package main

import (
	"fmt"
	"game-server-golang/cmd"
	"game-server-golang/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	err = cmd.ExecuteApi(config)
	if err != nil {
		panic(err)
	}
}
