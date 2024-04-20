package main

import "fmt"

func commandHelp(c *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cl := range commandList() {
		fmt.Printf("%v: %v\n", cl.Name, cl.Description)
	}
	fmt.Println()
	return nil
}
