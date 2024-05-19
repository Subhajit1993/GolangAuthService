package main

import (
	"Authentication/internal"
	"Authentication/internal/config"
	authenticator "Authentication/internal/config/authenticators"
	"Authentication/internal/config/database"
	"flag"
)

func main() {
	env := flag.String("env", "development", "environment to run the application in")
	flag.Parse()
	config.SetDirectoryPath("")
	config.LoadEnvironment(env)
	authenticator.InitAuth0()
	authenticator.InitWebAuthn()
	r := internal.RegisterRoutes()
	err := database.InitPgDatabase()
	if err != nil {
		panic(err.Error())
	}
	r.StartServer()
}
