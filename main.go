package main

import (
	"GoSolusVMClientApi/command"
	"GoSolusVMClientApi/config"
	"GoSolusVMClientApi/entity"
	"os"
)

func main() {
	// Config
	var configuration config.JSON
	configuration.Load()

	// Args - action must be passed in args
	Action := os.Args[1]

	// Execution
	Response := command.Execute(Action, entity.JSON(configuration))
	command.Output(Response, Action)

	// Exit
	os.Exit(0)
}
