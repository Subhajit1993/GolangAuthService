package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvironment(env *string) {
	envFile := fmt.Sprintf(".env.%s", *env)
	err := os.Setenv("env", *env)
	if err != nil {
		return
	}
	err = godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading " + envFile + " file")
		os.Exit(1)
	}
	fmt.Printf("Environment loaded from %s\n", envFile)
}
