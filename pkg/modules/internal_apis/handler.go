package internal_apis

import "github.com/gin-gonic/gin"

func ValidateAuth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Login",
	})
}
