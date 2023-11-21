package main

import (
	"time"

	"github.com/JAgo2000/Pokedex/internal/pokeapi"
)

type config struct {
	nextLocationArea     *string
	previousLocationArea *string
	pokeapiClient        pokeapi.Client
	coughtPokemon        map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		coughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	ReplLoop(&cfg)

}
