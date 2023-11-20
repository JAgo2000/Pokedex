package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationArea)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" -%v", area.Name)
		fmt.Println()
	}
	cfg.nextLocationArea = resp.Next
	cfg.previousLocationArea = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationArea == nil {
		return errors.New("you in the beginning")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationArea)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" -%v", area.Name)
		fmt.Println()
	}
	cfg.nextLocationArea = resp.Next
	cfg.previousLocationArea = resp.Previous
	return nil
}
