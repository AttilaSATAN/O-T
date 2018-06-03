package pokedex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var data BaseData
var pokemonIndexByTypeName map[string]int

func createIndexes() error {

	for n, p := range data.Pokemons {
		fmt.Println(n, p)
		PokemonIndexByName[p.Name] = n

		pok := &data.Pokemons[n]
		pok.IndexID, _ = strconv.Atoi(p.Number)
		pok.IndexName = []string{strings.ToLower(pok.Name)}
		pok.IndexTypeI = []string{}
		for _, t1 := range p.TypeI {
			pok.IndexTypeI = append(pok.IndexTypeI, t1)
		}
	}

	return nil
}

// populate read data.json and unmarshals it to a BaseData instance
func populate() error {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return err
	}

	json.Unmarshal(raw, &data)

	return createIndexes()
}

// New creates a new `PokedexServer` instance registers handlers and populate virtual database
func New() (*PokedexServer, error) {

	if err := populate(); err != nil {
		return nil, err
	}

	server := NewServer(8080)
	server.AddEndpoint("/", otherwiseHandler)
	server.AddEndpoint("/api/pokemon/list", listHandler)
	server.AddEndpoint("/api/pokemon/list/all", listHandler)
	return server, nil
}
