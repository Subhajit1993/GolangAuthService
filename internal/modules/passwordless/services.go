package passwordless

import (
	authenticator "Authentication/internal/config/authenticators"
	db "Authentication/internal/config/database"
	"Authentication/internal/entities"
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

func (request WebAuthNCredentialCreation) saveData(profile PublicProfile) error {
	userHandleString := ((request.Response.User.ID.(interface{})).(protocol.URLEncodedBase64)).String()
	passwordless := &entities.Passwordless{
		UserId:          profile.ID,
		UserHandle:      userHandleString,
		RPID:            request.Response.RelyingParty.ID,
		Challenge:       request.Response.Challenge.String(),
		CredentialType:  "platform",
		Counter:         0,
		AttestationData: string(request.Response.Attestation),
	}
	err := passwordless.Validate()
	if err != nil {
		return err
	}
	if err := db.PgDB.Create(passwordless).Error; err != nil {
		return err
	}
	return nil
}
