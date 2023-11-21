package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error { //variadic string parameter => (args ...string)
	if len(args) != 1 {
		return errors.New("no PokemonName provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.coughtPokemon[pokemonName]

	if !ok {
		return errors.New("dont catched yet!")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("------------\n")
	fmt.Printf("Types: \n")
	for _, typ := range pokemon.Types {
		fmt.Printf("- %v\n", typ.Type.Name)
	}
	fmt.Printf("------------\n")
	return nil
}
