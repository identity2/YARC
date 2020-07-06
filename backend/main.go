package main

import (
	"os"

	"github.com/YuChaoGithub/YARC/backend/app"
	"github.com/YuChaoGithub/YARC/backend/config"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")

	config := config.GetConfig()

	app := &app.App{}
	app.InitializeAndRun(config, jwtSecret)
}
