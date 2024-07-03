package passwordless

import (
	authenticator "Authentication/pkg/config/authenticators"
	db "Authentication/pkg/config/database"
	"Authentication/pkg/entities"
	"encoding/json"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"strconv"
)

type WebAuthNCredentialCreation struct {
	protocol.CredentialCreation
	webauthn.SessionData
}

func (request PublicProfile) PasswordlessRegistrationBeginAPI() WebAuthNCredentialCreation {
	user := authenticator.User{
		Id:          []byte(strconv.Itoa(request.ID)),
		Name:        request.FullName,
		DisplayName: request.DisplayName,
	}
	webAuthRegData, sessionData, err := authenticator.BeginRegistration(user)
	if err != nil {
		panic(err)
	}
	return WebAuthNCredentialCreation{
		CredentialCreation: *webAuthRegData,
		SessionData:        *sessionData,
	}
}

func (request WebAuthNCredentialCreation) saveData(profile PublicProfile) error {
	rawAttestationObject, err := json.Marshal(request.Response)
	sessionDataObject, err := json.Marshal(request.SessionData)
	userId, err := json.Marshal(request.SessionData.UserID)
	if err != nil {
		return err
	}
	rawAttestationObjectString := string(rawAttestationObject)
	sessionDataString := string(sessionDataObject)
	userIdStr := string(userId)

	passwordless := &entities.Passwordless{
		UserId:                         profile.ID,
		UserHandle:                     userIdStr,
		RPID:                           request.Response.RelyingParty.ID,
		Credential:                     request.Challenge,
		CredentialType:                 "platform",
		Status:                         entities.REGISTRATION_IN_PROGRESS,
		Counter:                        0,
		ExpiredAt:                      request.Expires,
		RawRegistrationAttestationData: rawAttestationObjectString,
		RawSessionData:                 sessionDataString,
		Active:                         false,
		IsDeleted:                      false,
	}
	err = passwordless.Validate()
	if err != nil {
		return err
	}
	if err := db.PgDB.Create(passwordless).Error; err != nil {
		return err
	}
	return nil
}
