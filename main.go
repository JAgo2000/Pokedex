package main

import (
	"time"

	"github.com/JAgo2000/Pokedex/internal/pokeapi"
)

type config struct {
	nextLocationArea     *string
	previousLocationArea *string
	pokeapiClient        pokeapi.Client
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	ReplLoop(&cfg)

}
