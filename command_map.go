package main

import (
	"fmt"

	"github.com/JAgo2000/Pokedex/internal/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" -%v", area.Name)
		fmt.Println()
	}
	return nil
}
