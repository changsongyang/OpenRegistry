package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/containerish/OpenRegistry/types"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (a *auth) RenewAccessToken(ctx echo.Context) error {
	ctx.Set(types.HandlerStartTime, time.Now())

	c, err := ctx.Cookie("refresh")
	if err != nil {
		if err == http.ErrNoCookie {
			echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error":   err.Error(),
				"message": "Unauthorised",
			})
			a.logger.Log(ctx, err)
			return echoErr
		}
		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "error getting refresh cookie",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	privBz, err := os.ReadFile(a.c.Registry.TLS.PrivateKey)
	if err != nil {
		panic(err)
	}
	privkey, err := jwt.ParseRSAPrivateKeyFromPEM(privBz)
	if err != nil {
		panic(err)
	}

	refreshCookie := c.Value
	var claims Claims
	tkn, err := jwt.ParseWithClaims(refreshCookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return privkey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
				"error":   err.Error(),
				"message": "signature error, unauthorised",
			})
			a.logger.Log(ctx, err)
			return echoErr
		}

		echoErr := ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   err.Error(),
			"message": "error parsing claims",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	if !tkn.Valid {
		err = fmt.Errorf("invalid token, Unauthorised")
		echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error":   err.Error(),
			"message": "invalid token, unauthorised",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	userId := claims.Id
	user, err := a.pgStore.GetUserById(ctx.Request().Context(), userId, false, nil)
	if err != nil {
		echoErr := ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error":   err.Error(),
			"message": "user not found in database, unauthorised",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	opts := &WebLoginJWTOptions{
		Id:        userId,
		Username:  user.Username,
		TokenType: "access_token",
		Audience:  a.c.Registry.FQDN,
		Privkey:   a.c.Registry.TLS.PrivateKey,
		Pubkey:    a.c.Registry.TLS.PubKey,
	}
	tokenString, err := NewWebLoginToken(opts)
	if err != nil {
		echoErr := ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error":   err.Error(),
			"message": "error creating new web token",
		})
		a.logger.Log(ctx, err)
		return echoErr
	}

	accessCookie := a.createCookie("access_token", tokenString, true, time.Now().Add(time.Hour))
	ctx.SetCookie(accessCookie)
	err = ctx.NoContent(http.StatusNoContent)
	a.logger.Log(ctx, err)
	return err
}
