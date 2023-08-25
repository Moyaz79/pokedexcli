package main

import (
	"fmt"
	"log"

	pokeapi "github.com/Moyaz79/pokedexcli/internal/pokeApi"
)

func main() {
	// startRepl()
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
