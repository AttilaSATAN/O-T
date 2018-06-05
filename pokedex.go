package pokedex

import (
	"encoding/json"
	"io/ioutil"
)

// populate read data.json and unmarshals it to a BaseData instance
func populate() error {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return err
	}

	json.Unmarshal(raw, &data)

	/*
	* Sıralama, metin manimülasyonu ve filtreleme gibi metodlarla BaseData'yı
	* dolaşarak sonuç toplayabilirdik ancak bu uzun vadede işlem gücü ihtiyacını arttıracaktır.
	* Bunun yerine indeksleme yaparak bir kaç kilobyte rem ile microsunucudan tasarruf sağlayabiliriz.
	 */
	return createIndexes()
}

// New creates a new `Server` instance registers handlers and populate virtual database
func New() (*Server, error) {

	if err := populate(); err != nil {
		return nil, err
	}

	server = NewServer(8080)
	server.AddEndpoint("/", otherwiseHandler)
	server.AddEndpoint("/api/pokemon/list", pokemonListHandler)
	server.AddEndpoint("/api/type/list", typeListHandler)
	server.AddEndpoint("/api/move/list", moveListHandler)
	for pokName, pok := range PokemonIndexByName {
		server.AddEndpoint("/api/pokemon/"+pokName, pokemonByNameHandlerProvider(*pok))
	}

	return server, nil
}
