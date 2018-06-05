package pokedex

import (
	"strconv"
	"strings"
)

//Pokemon represents a pokemon
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
	IndexID             int
	Indexes             map[string]map[string]bool
}

func createIndexes() error {

	TypeIndexByName = map[string]*Type{}
	MoveIndexByName = map[string]*Move{}
	PokemonIndexByName = map[string]*Pokemon{}
	for nt, t := range data.Types {
		ty := &data.Types[nt]
		TypeIndexByName[t.Name] = ty
		TypeIndexByName[strings.ToLower(ty.Name)] = ty
		TypeIndexByName[strings.ToUpper(ty.Name)] = ty
	}

	for n, p := range data.Pokemons {

		pok := &data.Pokemons[n]
		PokemonIndexByName[p.Name] = pok
		PokemonIndexByName[strings.ToLower(pok.Name)] = pok
		PokemonIndexByName[strings.ToUpper(pok.Name)] = pok

		pok.IndexID, _ = strconv.Atoi(p.Number)
		pok.Indexes = make(map[string]map[string]bool)

		pok.Indexes["name"] = make(map[string]bool)
		pok.Indexes["name"][strings.ToLower(pok.Name)] = true

		pok.Indexes["typeI"] = make(map[string]bool)
		for _, t1 := range p.TypeI {
			pok.Indexes["typeI"][strings.ToLower(t1)] = true
		}

		pok.Indexes["typeII"] = make(map[string]bool)
		for _, t2 := range p.TypeI {
			pok.Indexes["typeII"][strings.ToLower(t2)] = true
		}
		/*
			.
			.
			.
			ARTTIRILABİLİR
		*/

	}

	return nil
}

//IsIn
func (p *Pokemon) IsIn(index string, value string) bool {

	if _, exists := p.Indexes[index]; !exists {
		return false
	}

	_, exists := p.Indexes[index][strings.ToLower(value)]

	return exists
}
