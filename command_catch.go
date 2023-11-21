package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error { //variadic string parameter => (args ...string)
	if len(args) != 1 {
		return errors.New("no PokemonName provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience) //generates a number between 0 and (pokemon.BaseExperience)
	fmt.Printf("Baseexpirence %v,,,,Rand number%v\n", pokemon.BaseExperience, randNum)
	if randNum > threshold {
		return fmt.Errorf("failed to catch -%v\n", pokemonName)
	}
	cfg.coughtPokemon[pokemonName] = pokemon
	fmt.Printf("Success you catched -%v\n", pokemonName)
	return nil
}
