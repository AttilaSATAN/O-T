package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// notFoundHandler is 404 handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	written, err := json.Marshal(struct{ Message string }{Message: "Bulamadık :("})
	if err != nil {
		badImplementationHandler(w, r, err)
	}
	_, err = w.Write(written)
	if err != nil {

		badImplementationHandler(w, r, err)
	}
}

// badRequestHandler is general 4xx handler
func badRequestHandler(w http.ResponseWriter, r *http.Request, e error) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, e.Error())
}

// badImplementationHandler is general 5xx handler
func badImplementationHandler(w http.ResponseWriter, r *http.Request, e error) {
	log.Fatalln(e)
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Bir sorun oluştu. :(")
}

// otherwiseHandler is siple router for home and 404
func otherwiseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
	} else {
		wellcomeHandler(w, r)
	}
}

// wellcomeHandler is home handler
func wellcomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hoşgeldin :)")
}

//
func all(w http.ResponseWriter, r *http.Request) {

	pokemons, err := json.Marshal(data.Pokemons)
	if err != nil {
		badImplementationHandler(w, r, err)
	}
	_, err = w.Write(pokemons)
	if err != nil {
		badImplementationHandler(w, r, err)
	}
}

// pokemonListHandler for listin pokemons
func listHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	if val, pres := query["name"]; pres {
		fmt.Println(val)
		fmt.Println(query.Get("a"))
		//index := PokemonIndexByName[val]
	} else {
		all(w, r)
	}
}

func queryRouter(w http.ResponseWriter, r *http.Request) (bool, error) {

	return true, nil
}
