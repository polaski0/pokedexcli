package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := commandList()

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			return
		}

		t := scanner.Text()
		if c, ok := cmd[t]; ok {
			err := c.callback()

			if err != nil {
				return
			}
		}
	}
}

func commandList() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes the program",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
    fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
    fmt.Println()
	for _, c := range commandList() {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
