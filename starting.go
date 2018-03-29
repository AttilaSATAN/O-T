package main

import (
	"fmt"
	"log"
	"net/http"
)

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy          struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
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

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/list url:", r.URL)
	fmt.Fprint(w, "The List Handler\n")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/get url:", r.URL)
	fmt.Fprint(w, "The Get Handler\n")
}

func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	//TODO: read data.json to a BaseData

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/get", getHandler)
	//TODO: add more
	http.HandleFunc("/", otherwise)
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
