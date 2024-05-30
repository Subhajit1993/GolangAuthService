package passwordless

import (
	authenticator "Authentication/internal/config/authenticators"
	"github.com/gin-gonic/gin"
	"net/http"
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
	webAuthRegData.saveData(user)
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
}

func FinishRegistration(ctx *gin.Context) {
	webAuthRegData := authenticator.FinishRegistration(ctx.Request)
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthRegData})
}

func LoginBegin(ctx *gin.Context) {
	webAuthLoginData := authenticator.BeginLogin()
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": webAuthLoginData})
}
