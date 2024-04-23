package main

import "fmt"

func commandMap(c *Config, args ...string) error {
	locationRes, err := c.Client.GetLocation(c.Next)

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

func commandMapb(c *Config, args ...string) error {
	locationRes, err := c.Client.GetLocation(c.Previous)

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
