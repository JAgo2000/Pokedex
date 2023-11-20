package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error { //variadic string parameter => (args ...string)
	if len(args) != 1 {
		return errors.New("No location erea provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" -%v", pokemon.Pokemon.Name)
		fmt.Println()
	}
	return nil
}
