package main

import (
	"fmt"

	"github.com/polaski0/pokedexcli/internal/api"
)

func commandMap(c *Config) error {
	locationRes, err := api.GetLocationArea(c.Next)

	if err != nil {
		return err
	}

	c.Next = locationRes.Next
	c.Previous = locationRes.Previous

	for _, res := range locationRes.Results {
		fmt.Println(res.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	locationRes, err := api.GetLocationArea(c.Previous)

	if err != nil {
		return err
	}

	c.Next = locationRes.Next
	c.Previous = locationRes.Previous

	for _, res := range locationRes.Results {
		fmt.Println(res.Name)
	}

	return nil
}