package passwordless

import (
	authenticator "Authentication/pkg/config/authenticators"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
	"strconv"
)

func BeginRegistration(ctx *gin.Context) {
	userIdStr := ctx.MustGet("user_id").(float64)
	user := PublicProfile{ID: int(userIdStr)}
	user, err := user.findWithID()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	webAuthRegData := user.PasswordlessRegistrationBeginAPI()
	err = webAuthRegData.saveData(user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
}

func FinishRegistration(ctx *gin.Context) {
	userIdStr := ctx.MustGet("user_id").(float64)
	passwordlessData := passwordlessRegistration{UserId: int(userIdStr)}
	passwordlessData, err := passwordlessData.getPasswordlessRegistrationData()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	userIdBytes := []byte(strconv.Itoa(int(userIdStr)))

	webAuthNSessionData := webauthn.SessionData{
		UserID:           userIdBytes,
		Challenge:        passwordlessData.Challenge,
		UserVerification: "preferred",
		Expires:          passwordlessData.ExpiredAt,
	}

	webAuthNUser := authenticator.User{
		Id:          userIdBytes,
		Name:        passwordlessData.Name,
		DisplayName: passwordlessData.DisplayName,
	}

	webAuthRegData := authenticator.FinishRegistration(ctx.Request, &webAuthNSessionData, &webAuthNUser)
	// Save the credential ID to the database
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
	return
}

// LoginBegin initiates the passwordless login process by obtaining login data from the authenticator package and returning it in a JSON response with an "ok" status.
func LoginBegin(ctx *gin.Context) {
	webAuthLoginData := authenticator.BeginLogin()
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthLoginData})
}

// LoginBeginNew initiates a passwordless login process using a custom parameter. It calls authenticator.BeginLogin with "Hello" and responds with an HTTP 200 OK JSON object containing the login data.
func LoginBeginNew(ctx *gin.Context) {
	webAuthLoginData := authenticator.BeginLogin("Hello")
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthLoginData})
}
