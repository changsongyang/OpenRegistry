package webauthn

import (
	"github.com/containerish/OpenRegistry/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
)

type (
	WebAuthNSessiondata struct {
		webauthn.SessionData
		CredentialOwnerId string
	}

	WebAuthnUser struct {
		*types.User
		credentials []webauthn.Credential
	}
)

// WebAuthnID - User ID according to the Relying Party
// TODO(jay-dee7): This will panic if the uuid is not in the requited format
func (u *WebAuthnUser) WebAuthnID() []byte {
	userID := uuid.MustParse(u.Id)
	return userID[:]
}

// WebAuthnName - User Name according to the Relying Party
func (u *WebAuthnUser) WebAuthnName() string {
	return u.Username
}

// WebAuthnDisplayName - Display Name of the user
func (u *WebAuthnUser) WebAuthnDisplayName() string {
	return u.Username
}

// WebAuthnIcon - User's icon url
func (u *WebAuthnUser) WebAuthnIcon() string {
	return u.AvatarURL
}

// WebAuthnCredentials - Credentials owned by the user
func (u *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

func (u *WebAuthnUser) AddWebAuthNCredential(creds *webauthn.Credential) {
	// initialised to non-nil value in case of first attempt
	if u.credentials == nil {
		u.credentials = make([]webauthn.Credential, 0)
	}
	u.credentials = append(u.credentials, *creds)
}

func (u *WebAuthnUser) AddWebAuthNCredentials(creds ...*webauthn.Credential) {
	for _, c := range creds {
		if c == nil {
			continue
		}

		u.credentials = append(u.credentials, *c)
	}
}

func (u *WebAuthnUser) GetWebauthnCredentialDescriptors() []protocol.CredentialDescriptor {
	var list []protocol.CredentialDescriptor

	for _, cred := range u.credentials {
		list = append(list, protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		})
	}

	return list
}

func (u *WebAuthnUser) GetUnderlyingUser() *types.User {
	return u.User
}

func (u *WebAuthnUser) FromUnderlyingUser(user *types.User) *WebAuthnUser {
	return &WebAuthnUser{
		User: user,
	}
}
