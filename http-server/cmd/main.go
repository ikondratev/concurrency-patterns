package main

import (
	"fmt"
	"os"

	"github.com/kilia/http-example/internal/app"
)

const devEnv = "dev"

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = devEnv
	}

	application, err := app.New(env)
	if err != nil {
		fmt.Println("New application erro:", err)
		os.Exit(1)
	}

	if err := application.Run(); err != nil {
		fmt.Println("Server error:", err)
		os.Exit(1)
	}
}