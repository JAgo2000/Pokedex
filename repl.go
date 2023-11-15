package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints the possible commands",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "turns off the pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "list some locationAreas",
			callback:    callbackMap,
		},
	}
}

func ReplLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()
		text_arr := CleanScantext(text)
		if nil == text_arr {
			continue
		}
		commandName := text_arr[0]
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()

	}
}

func CleanScantext(text string) []string {
	if len(text) == 0 {
		return nil
	}
	text_lower := strings.ToLower(text)
	arr_str := strings.Fields(text_lower)
	return arr_str
}
