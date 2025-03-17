package pkg

import (
	"Authentication/pkg/middlewares"
	"Authentication/pkg/modules"
	"Authentication/pkg/modules/general"
	"Authentication/pkg/modules/internal_apis"
	"Authentication/pkg/modules/openid"
	"Authentication/pkg/modules/passwordless"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"log"
)
import "github.com/gin-gonic/gin"

type GinEngine struct {
	*gin.Engine
}

func setupGin() GinEngine {
	r := gin.Default()
	return GinEngine{Engine: r}
}

func (r GinEngine) addRoutes() GinEngine {
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	r.GET("/", modules.HealthCheck)
	r.GET("/health", modules.HealthCheck)
	r.Use(sessions.Sessions("auth-session", store))
	r.GET("/refresh-access-token", openid.GetAccessTokenRefreshToken)
	r.POST("/email/login", general.Login)
	r.GET("/logout", openid.Logout)
	devToolsApis := r.Group("/dev-tools")
	{
		devToolsApis.GET("/", openid.Home)
		devToolsApis.GET("/login", openid.Login)
		devToolsApis.GET("/callback", openid.Callback)
		devToolsApis.GET("/user", middlewares.ValidateJWTToken, openid.User)
		devToolsPasswordLessApis := devToolsApis.Group("/passwordless")
		{
			devToolsPasswordLessApis.POST("/begin-registration", middlewares.ValidateJWTToken, passwordless.BeginRegistration)
			devToolsPasswordLessApis.POST("/finish-registration", middlewares.ValidateJWTToken, passwordless.FinishRegistration)
			devToolsPasswordLessApis.POST("/begin-login", passwordless.LoginBegin)
		}
	}
	internalApis := r.Group("/internal")
	{
		internalApis.GET("/validate/auth", internal_apis.ValidateAuth)
	}
	return r
}

func (r GinEngine) addMiddleware() GinEngine {
	//r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())
	return r
}

func RegisterRoutes() GinEngine {

	r := setupGin()
	// global middleware
	r.Use(middlewares.CORSMiddleware())
	r.LoadHTMLGlob("web/template/*")
	r = r.addMiddleware()
	r.addRoutes()
	return r
}

func (r GinEngine) StartServer() {
	err := r.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
