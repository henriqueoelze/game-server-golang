package main

import "game-server-golang/cmd"

func main() {
	err := cmd.ExecuteApi()
	if err != nil {
		panic(err)
	}
}
