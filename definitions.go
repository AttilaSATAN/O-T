package pokedex

import (
	"sort"
)

//Type is pokemon type. Also it's a very ugly name for a struct.
type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

// Move is an attack information. The
type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

//PokemonList type for collecting pokemons.
type PokemonList []Pokemon

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type      `json:"types"`
	Pokemons PokemonList `json:"pokemons"`
	Moves    []Move      `json:"moves"`
}

// FilterByName is Array#filter translation for PokemonList
func (data PokemonList) FilterBy(index string, name string) PokemonList {
	ret := PokemonList{}
	for _, p := range data {
		if p.IsIn(index, name) {
			ret = append(ret, p)
		}
	}
	return ret
}

func (data PokemonList) Sort(by string) PokemonList {

	if by == "name" || by == "Name" || by == "NAME" {
		sort.Slice(data, func(i, j int) bool { return data[i].Name < data[j].Name })
	}

	if by == "baseAttack" || by == "Base Attack" || by == "BASE ATTACK" || by == "BaseAttack" || by == "baseattack" {
		sort.Slice(data, func(i, j int) bool { return data[i].BaseAttack < data[j].BaseAttack })
	}
	if by == "-name" || by == "-Name" || by == "-NAME" {
		sort.Slice(data, func(i, j int) bool { return data[i].Name > data[j].Name })
	}

	if by == "-baseAttack" || by == "-Base Attack" || by == "-BASE ATTACK" || by == "-BaseAttack" || by == "-baseattack" {
		sort.Slice(data, func(i, j int) bool { return data[i].BaseAttack > data[j].BaseAttack })
	}
	/*
	   .
	   .
	   .
	*/
	return data

}

//data is for storage
var data BaseData

// PokemonIndexByName is a map for pokemons by their name as keys
var PokemonIndexByName map[string]*Pokemon
var TypeIndexByName map[string]*Type
var MoveIndexByName map[string]*Move
var server *Server
