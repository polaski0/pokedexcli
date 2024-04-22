package main

import (
	"fmt"

	"github.com/polaski0/pokedexcli/internal/api"
)

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	area := args[0]
	exploredRes, err := api.GetExploredArea(&area)

	if err != nil {
		return err
	}

    fmt.Printf("Exploring %v...\n", area)
	for _, res := range exploredRes.PokemonEncounters {
		fmt.Println("-", res.Pokemon.Name)
	}

	return nil
}
