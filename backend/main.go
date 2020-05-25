package main

import (
	"github.com/YuChaoGithub/YARC/backend/app"
	"github.com/YuChaoGithub/YARC/backend/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.InitializeAndRun(config)
}
