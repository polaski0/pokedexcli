package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/polaski0/pokedexcli/internal/api"
)

type Config struct {
	Client   api.Client
	Pokedex  map[string]api.Pokemon
	Next     *string
	Previous *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config, args ...string) error
}

func startRepl(conf *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := commandList()

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			return
		}

		t := cleanInput(scanner.Text())
		commandName := t[0]

		if c, ok := cmd[commandName]; ok {
			err := c.Callback(conf, t[1:]...)

			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func cleanInput(s string) []string {
	out := strings.ToLower(s)
	w := strings.Fields(out)
	return w
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
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists all caught pokemons",
			Callback:    commandPokedex,
		},
	}
}
