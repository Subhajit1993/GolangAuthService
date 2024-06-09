package authenticator

import (
	"encoding/json"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id          []byte
	Name        string
	DisplayName string
	Icon        string
	credentials []webauthn.Credential
}

func (u *User) WebAuthnID() []byte {
	return u.Id
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnIcon() string {
	return u.Icon
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

var (
	webAuthn *webauthn.WebAuthn
)

func InitWebAuthn() {
	webAuthn, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Example Inc.",
		RPID:          "localhost",
		RPOrigin:      "http://localhost:8080",
	})
	if err != nil {
		panic(err)
	}
}

func BeginRegistration(u User) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	user := &u
	options, session, err := webAuthn.BeginRegistration(user)
	if err != nil {
		return nil, nil, err
	}
	// Store sessionData somewhere, e.g. in the user's session
	// You'll need this in the next step of the registration process
	// store it in a temp file for now
	sessionDataJSON, err := json.Marshal(session)
	err = os.WriteFile("sessionData", sessionDataJSON, 0644)
	if err != nil {
		return nil, nil, err
	}
	return options, session, nil
}

type CreationResponses protocol.CredentialCreationResponse

func FinishRegistration(r *http.Request, sessionData *webauthn.SessionData, webAuthNUser *User) *webauthn.Credential {
	// Read the session data from the file
	/*sessionDataJSON, err := os.ReadFile("sessionData")
	if err != nil {
		log.Println("Failed to read session data:", err)
		return false
	}
	session := new(webauthn.SessionData)
	err = json.Unmarshal(sessionDataJSON, session)
	if err != nil {
		log.Println("Failed to decode session data:", err)
		return false
	}

	// Create a user
	//user := &User{Id: []byte("user-id"), Name: "subhajitd@plateron.com", DisplayName: "Subhajit Dutta"}*/
	credential, err := webAuthn.FinishRegistration(webAuthNUser, *sessionData, r)
	if err != nil {
		log.Println("Failed to finish registration:", err)
		return nil
	}

	// Write the credential data to a file
	credentialJSON, err := json.Marshal(credential)
	if err != nil {
		log.Println("Failed to encode credential data:", err)
		return nil
	}
	err = os.WriteFile("credential", credentialJSON, 0644)
	if err != nil {
		log.Println("Failed to write credential data:", err)
		return nil
	}

	return credential
}

func BeginLogin() *protocol.CredentialAssertion {
	// Read the credential data from the file
	credentialJSON, err := os.ReadFile("credential")
	if err != nil {
		log.Println("Failed to read credential data:", err)
		return nil
	}
	credential := new(webauthn.Credential)
	err = json.Unmarshal(credentialJSON, credential)
	user := &User{Id: []byte("user-id"),
		Name:        "subhajitd@plateron.com",
		DisplayName: "Subhajit Dutta",
		credentials: []webauthn.Credential{*credential}}
	options, session, err := webAuthn.BeginLogin(user)
	if err != nil {
		return nil
	}
	// Store sessionData somewhere, e.g. in the user's session
	// You'll need this in the next step of the login process
	// store it in a temp file for now
	sessionDataJSON, err := json.Marshal(session)
	err = os.WriteFile("sessionData", sessionDataJSON, 0644)
	if err != nil {
		return nil
	}
	return options
}
