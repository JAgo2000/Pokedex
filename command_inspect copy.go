package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error { //variadic string parameter => (args ...string)
	fmt.Println("Pokemon in Pokedex:")

	for _, pokemon := range cfg.coughtPokemon {
		fmt.Printf("- %v\n", pokemon.Name)
	}

	return nil
}
