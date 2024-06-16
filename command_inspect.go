package main

import "fmt"

func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return nil
	}

	pokemonName := args[0]
	if pokemon, ok := c.Pokedex[pokemonName]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, s := range pokemon.Stats {
			fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %v\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
