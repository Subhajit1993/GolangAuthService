package general

import (
	authenticator "Authentication/pkg/config/authenticators"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	var req LoginRequest
	session := sessions.Default(ctx)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user, err := req.findWithEmail()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error while finding user",
		})
		return
	}
	req.ID = user.ID
	jwtToken, err := authenticator.CreateToken(user.ID)
	refreshToken, expiry, err := authenticator.CreateRefreshToken(user.ID, 48)
	err = req.saveRefreshToken(refreshToken, expiry)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", jwtToken)
	session.Set("refresh_token", refreshToken)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
