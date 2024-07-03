package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func SetDirectoryPath(path string) {
	err := os.Setenv("root_path", path)
	if err != nil {
		return
	}
}

func LoadEnvironment() {
	env := os.Getenv("GOENV")
	envFile := fmt.Sprintf(".env.%s", env)
	err := os.Setenv("env", env)
	fmt.Println("Environment: ", env)
	if err != nil {
		return
	}
	envFile = os.Getenv("root_path") + envFile
	err = godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading " + envFile + " file")
		os.Exit(1)
	}
	fmt.Printf("Environment loaded from %s\n", envFile)
}
