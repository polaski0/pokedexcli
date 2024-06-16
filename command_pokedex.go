package main

import "fmt"

func commandPokedex(c *Config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for pokemon := range c.Pokedex {
		fmt.Printf("- %v\n", pokemon)
	}

	return nil
}
