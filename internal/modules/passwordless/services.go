package passwordless

import (
	authenticator "Authentication/internal/config/authenticators"
	"github.com/go-webauthn/webauthn/protocol"
)

func (request PasswordlessRegistrationBeginAPIRequest) PasswordlessRegistrationBeginAPI() *protocol.CredentialCreation {
	user := authenticator.User{
		Id:          []byte(request.Email),
		Name:        request.Email,
		DisplayName: request.DisplayName,
	}
	webAuthRegData := authenticator.BeginRegistration(user)
	return webAuthRegData
}
