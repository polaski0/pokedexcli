package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	pokemon := args[0]
	pokemonRes, err := c.Client.GetPokemon(&pokemon)

	if err != nil {
		return err
	}

    fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
    isCaught := math.Round(float64(rand.Intn(pokemonRes.BaseExperience)) / float64(pokemonRes.BaseExperience))
	if isCaught == 1 {
		fmt.Printf("%v was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
		c.Pokedex[pokemon] = pokemonRes
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}

	return nil
}
