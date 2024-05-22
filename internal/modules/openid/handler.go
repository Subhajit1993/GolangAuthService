package openid

import (
	authenticator "Authentication/internal/config/authenticators"
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func Login(ctx *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Save the state inside the session.
	session := sessions.Default(ctx)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, authenticator.Auth.AuthCodeURL(state))
}

func Callback(ctx *gin.Context) {
	session := sessions.Default(ctx)
	state := session.Get("state")
	if state != ctx.Query("state") {
		ctx.String(http.StatusUnauthorized, "Invalid state")
		return
	}

	token, err := authenticator.Auth.Exchange(ctx, ctx.Query("code"))
	if err != nil {
		ctx.String(http.StatusUnauthorized, err.Error())
		return
	}

	idToken, err := authenticator.Auth.VerifyIDToken(ctx.Request.Context(), token)
	if err != nil {
		ctx.String(http.StatusUnauthorized, err.Error())
		return
	}

	var profile Profile
	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	profile.Email = claims["email"].(string)

	existingUser, err := profile.findWithEmail()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if existingUser != nil {
		session.Set("access_token", token.AccessToken)
		session.Set("profile", claims)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Redirect(http.StatusTemporaryRedirect, "/dev-tools/user")
		return
	}

	profile.FullName = claims["name"].(string)
	profile.DisplayName = claims["nickname"].(string)
	profile.RegistrationSource = claims["sub"].(string)
	profile.Verified = claims["email_verified"].(bool)
	profile.Picture = claims["picture"].(string)

	_, err = profile.saveData()
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", claims)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, "/dev-tools/user")
}

func User(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	log.Println(profile)

	ctx.HTML(http.StatusOK, "user.html", profile)
}

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}

func Logout(ctx *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()
	session := sessions.Default(ctx)
	session.Clear()
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
