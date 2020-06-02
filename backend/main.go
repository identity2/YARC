package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/YuChaoGithub/YARC/backend/app"
	"github.com/YuChaoGithub/YARC/backend/config"
)

func main() {
	fmt.Printf("JWT Secret Key...")
	jwtSecret, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error: Failed to get the JWT secret key.")
		return
	}

	config := config.GetConfig()

	app := &app.App{}
	app.InitializeAndRun(config, jwtSecret)
}
