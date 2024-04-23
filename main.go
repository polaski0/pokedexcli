package main

import (
	"time"

	"github.com/polaski0/pokedexcli/internal/api"
)

func main() {
    c := api.NewClient(30 * time.Second)
    conf := &Config{
        Client: c,
    }

	startRepl(conf)
}
