package main

import (
	"Authentication/pkg"
	"Authentication/pkg/config"
	authenticator "Authentication/pkg/config/authenticators"
	"Authentication/pkg/config/database"
)

func main() {
	//env := flag.String("env", "development", "environment to run the application in")
	//flag.Parse()
	//config.SetDirectoryPath("")
	config.LoadEnvironment()
	authenticator.InitAuth0()
	authenticator.InitWebAuthn()
	r := pkg.RegisterRoutes()
	err := database.InitPgDatabase()
	if err != nil {
		panic(err.Error())
	}
	r.StartServer()
}
