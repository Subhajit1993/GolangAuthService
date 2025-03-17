package main

import (
	"Authentication/pkg"
	"Authentication/pkg/config"
	authenticator "Authentication/pkg/config/authenticators"
	"Authentication/pkg/config/database"
)

// main is the entry point of the authentication application. It loads the environment configuration, initializes the authentication mechanisms (Auth0 and WebAuthn), registers the application routes, establishes the PostgreSQL database connection, and starts the HTTP server. The function panics if the database initialization fails.
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
