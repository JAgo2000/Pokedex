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
	callback    func(*config, ...string) error
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
		"mapb": {
			name:        "mapb",
			description: "list previus locationAreas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location Area}",
			description: "lists the pokemon in a location area and add it to your pokedex",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "try to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "try to inspect a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all Pokemons that you cought",
			callback:    callbackPokedex,
		},
	}
}

func ReplLoop(cfg *config) {
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
		args := []string{}
		if len(text_arr) > 1 {
			args = text_arr[1:]
		}
		commandName := text_arr[0]
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

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
