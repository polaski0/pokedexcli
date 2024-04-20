package main

import (
	"fmt"

	"github.com/polaski0/pokedexcli/internal"
)

func commandMap(c *Config) error {
    locationRes, err := internal.GetLocationArea(c.Next)

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
    locationRes, err := internal.GetLocationArea(c.Previous)

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
