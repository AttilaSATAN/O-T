package main

import "github.com/attilasatan/pokedex"

func main() {
	_, pokedexServer := pokedex.New()
	pokedexServer.Serve()
}
