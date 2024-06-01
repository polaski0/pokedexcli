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

	isCaught := math.Round(float64(rand.Intn(pokemonRes.BaseExperience)) / float64(pokemonRes.BaseExperience))

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	if isCaught == 1 {
		fmt.Printf("%v was caught!\n", pokemon)
		c.Client.Pokedex[pokemon] = pokemonRes
	} else {
		fmt.Printf("%v escaped!\n", pokemon)
	}

	return nil
}
