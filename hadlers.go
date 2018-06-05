package pokedex

import (
	"encoding/json"
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

func responder(w http.ResponseWriter, r *http.Request, data interface{}) (err error) {

	pokemons, err := json.Marshal(data)
	if err != nil {
		return
	}
	_, err = w.Write(pokemons)
	if err != nil {
		return
	}
	return

}

// pokemonListHandler for listin pokemons
func pokemonListHandler(w http.ResponseWriter, r *http.Request) {

	if responded, err := pokemonQueryRouter(w, r); responded {
		if err != nil {
			badImplementationHandler(w, r, err)
		}
	} else {
		responder(w, r, data.Pokemons)
	}
}

// pokemonListHandler for listing pokemon types
func typeListHandler(w http.ResponseWriter, r *http.Request) {

	responder(w, r, data.Types)

}

// pokemonListHandler for listing moves
func moveListHandler(w http.ResponseWriter, r *http.Request) {

	responder(w, r, data.Moves)

}

func pokemonQueryRouter(w http.ResponseWriter, r *http.Request) (bool, error) {

	pokemons := data.Pokemons
	query := r.URL.Query()

	for _, index := range []string{"name", "typeI", "typeII"} {
		if _, pres := query[index]; pres {
			name := query.Get(index)
			pokemons = pokemons.FilterBy(index, name)
		}
	}

	if _, s := query["sortBy"]; s {
		by := query.Get("sortBy")
		pokemons = pokemons.Sort(by)
	}

	if err := responder(w, r, pokemons); err != nil {
		return false, err
	}

	return true, nil
}

func pokemonByNameHandlerProvider(p Pokemon) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		pokemon, err := json.Marshal(p)
		if err != nil {
			return
		}
		_, err = w.Write(pokemon)
		if err != nil {
			return
		}
		return
	}
}
