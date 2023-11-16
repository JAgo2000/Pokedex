package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("*****************")
	fmt.Println("This is a Repl! ")
	fmt.Println("Valid Commands: ")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("<%v> %v", command.name, command.description)
		fmt.Println()
	}
	fmt.Println("*****************")
	return nil
}
