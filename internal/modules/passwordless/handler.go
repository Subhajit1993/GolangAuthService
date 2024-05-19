package passwordless

import (
	authenticator "Authentication/internal/config/authenticators"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PasswordlessBeginRegistration(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	if profile == nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	log.Println(profile)
	var registrationData PasswordlessRegistrationBeginAPIRequest
	err := ctx.ShouldBindJSON(&registrationData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	webAuthRegData := registrationData.PasswordlessRegistrationBeginAPI()
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
}

func PasswordLessFinishRegistration(ctx *gin.Context) {
	webAuthRegData := authenticator.FinishRegistration(ctx.Request)
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
}

func PasswordlessLoginBegin(ctx *gin.Context) {
	webAuthLoginData := authenticator.BeginLogin()
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthLoginData})
}
