package main

import (
	"awsc/logger"
	"os"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	logger.InitLogger(env)
	defer logger.Sync()
}
