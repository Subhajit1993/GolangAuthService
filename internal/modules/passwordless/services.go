package passwordless

import (
	authenticator "Authentication/internal/config/authenticators"
	"Authentication/internal/entities"
	"fmt"
	"github.com/go-webauthn/webauthn/protocol"
)

type WebAuthNCredentialCreation protocol.CredentialCreation

func (request PublicProfile) PasswordlessRegistrationBeginAPI() WebAuthNCredentialCreation {
	user := authenticator.User{
		Id:          []byte(request.Email),
		Name:        request.Email,
		DisplayName: request.DisplayName,
	}
	webAuthRegData := authenticator.BeginRegistration(user)
	return WebAuthNCredentialCreation(*webAuthRegData)
}

func (request WebAuthNCredentialCreation) saveData(profile PublicProfile) {
	passwordless := &entities.Passwordless{
		UserId:    profile.ID,
		Challenge: string(request.Response.Challenge),
	}
	fmt.Println(passwordless)
}
