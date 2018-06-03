package main

import "github.com/attilasatan/pokedex"

func main() {
	pokedexServer, err := pokedex.New()
	if err != nil {
		panic(err)
	}
	pokedexServer.Serve()
}
