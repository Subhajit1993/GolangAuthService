package middlewares

import (
	authenticator "Authentication/internal/config/authenticators"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticatedAuth0(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Next()
	}
}

func ValidateJWTToken(ctx *gin.Context) {
	session := sessions.Default(ctx)
	token := session.Get("access_token")

	if token == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	verifiedToken, err := authenticator.ValidateToken(token.(string))
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")

	}
	if verifiedToken == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// add user_id to context
	ctx.Set("user_id", verifiedToken.Claims.(jwt.MapClaims)["user_id"])
	ctx.Next()
}
