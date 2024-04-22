package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := commandList()
	var conf Config

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			return
		}

		t := strings.Split(scanner.Text(), " ")
		if c, ok := cmd[cleanInput(t[0])]; ok {
			err := c.Callback(&conf) // allow callback functions to retrieve other arguments

			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func cleanInput(s string) string {
    return strings.ToLower(s)
}

func commandList() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Closes the program",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "The map command displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "It displays the previous 20 locations. It's a way to go back.",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the location related to the area",
			Callback:    commandExplore,
		},
	}
}
