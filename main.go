package main

import (
	"Authentication/pkg"
	"Authentication/pkg/config"
	authenticator "Authentication/pkg/config/authenticators"
	"Authentication/pkg/config/database"
)

// main is the entry point for the authentication application. It loads the environment configuration,
// initializes authentication methods (Auth0 and WebAuthn), registers application routes, and sets up the
// PostgreSQL database connection, panicking if the connection fails, before starting the server.
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
	// Hello
	r.StartServer()
}
