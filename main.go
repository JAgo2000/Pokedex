package main

import (
	"fmt"
	"log"

	"github.com/JAgo2000/Pokedex/internal/pokeapi"
)

func main() {
	//ReplLoop() testing api
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err) //exits the program os.Exit(1)
	}
	fmt.Println(resp)

}
