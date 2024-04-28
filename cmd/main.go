package main

import (
	"Authentication/internal/config"
	"flag"
)

func main() {
	env := flag.String("env", "development", "environment to run the application in")
	flag.Parse()
	config.LoadEnvironment(env)
}
