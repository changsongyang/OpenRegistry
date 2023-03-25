package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/containerish/OpenRegistry/types"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

func (a *auth) SignIn(ctx echo.Context) error {
	ctx.Set(types.HandlerStartTime, time.Now())
	var user types.User

	if err := json.NewDecoder(ctx.Request().Body).Decode(&user); err != nil {
		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "invalid JSON object",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	err := user.Validate(true)
	if err != nil {
		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "invalid data provided for user login",
			"code":    "INVALID_CREDENTIALS",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	key := user.Email
	if user.Username != "" {
		key = user.Username
	}

	userFromDb, err := a.pgStore.GetUser(ctx.Request().Context(), key, true, nil)
	if err != nil {

		if errors.Unwrap(err) == pgx.ErrNoRows {
			echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
				"error":   err.Error(),
				"message": "user not found",
			})
			a.logger.Log(ctx, err)
			return echoErr
		}

		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "database error, failed to get user",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	if !userFromDb.IsActive {
		err = fmt.Errorf("account is inactive, please check your email and verify your account")
		echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error":   "ERR_USER_INACTIVE",
			"message": err.Error(),
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	if !a.verifyPassword(userFromDb.Password, user.Password) {
		err = fmt.Errorf("password is incorrect")
		echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error":   "ERR_INCORRECT_PASSWORD",
			"message": err.Error(),
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	access, err := a.newWebLoginToken(userFromDb.Id, userFromDb.Username, "access")
	if err != nil {
		echoErr := ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err.Error(),
			"message": "error creating web login token",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	refresh, err := a.newWebLoginToken(userFromDb.Id, userFromDb.Username, "refresh")
	if err != nil {
		echoErr := ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err.Error(),
			"message": "error creating refresh token",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	id, err := uuid.NewRandom()
	if err != nil {
		a.logger.Log(ctx, err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err.Error(),
			"message": "error creating session id",
		})
	}
	if err = a.pgStore.AddSession(ctx.Request().Context(), id.String(), refresh, userFromDb.Username); err != nil {
		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "error creating session",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	sessionId := fmt.Sprintf("%s:%s", id, userFromDb.Id)
	sessionCookie := a.createCookie("session_id", sessionId, false, time.Now().Add(time.Hour*750))
	accessCookie := a.createCookie("access", access, true, time.Now().Add(time.Hour*750))
	refreshCookie := a.createCookie("refresh", refresh, true, time.Now().Add(time.Hour*750))

	ctx.SetCookie(accessCookie)
	ctx.SetCookie(refreshCookie)
	ctx.SetCookie(sessionCookie)
	err = ctx.JSON(http.StatusOK, echo.Map{
		"token":   access,
		"refresh": refresh,
	})
	a.logger.Log(ctx, err)
	return err
}
