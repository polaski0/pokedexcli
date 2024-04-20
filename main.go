package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const API_BASE_PATH = "https://pokeapi.co/api/v2/"
const LOCATION_AREA_PATH = API_BASE_PATH + "location-area"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

// Convert Next and Previous to a nullable string, turning string into *string
type Config struct {
	Next     string
	Previous string
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmd := commandList()
	conf := Config{
		Next:     LOCATION_AREA_PATH,
		Previous: "",
	}

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()

		if !ok {
			return
		}

		t := scanner.Text()
		if c, ok := cmd[t]; ok {
			err := c.Callback(&conf)

			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

// Below this line are different commands and its callback functions
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
	}
}

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

func commandExit(c *Config) error {
	os.Exit(0)
	return nil
}

func getLocationArea(c *Config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	var jsonRes LocationArea
	err = json.Unmarshal(body, &jsonRes)

	if err != nil {
		return err
	}

	c.Next = jsonRes.Next
	c.Previous = jsonRes.Previous

	for _, v := range jsonRes.Results {
		fmt.Println(v.Name)
	}

    return nil
}

func commandMap(c *Config) error {
    if c.Next == "" {
        fmt.Println("No more area to explore...")
        return nil
    }

    err := getLocationArea(c, c.Next)
    if err != nil {
        return err
    }

	return nil
}

func commandMapb(c *Config) error {
    if c.Previous == "" {
        fmt.Println("No more area to explore...")
        return nil
    }

    err := getLocationArea(c, c.Previous)
    if err != nil {
        return err
    }

	return nil
}
