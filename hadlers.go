package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func badRequestHandler(w http.ResponseWriter, r *http.Request, e error) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, e.Error())
}

func badImplementationHandler(w http.ResponseWriter, r *http.Request, e error) {
	log.Fatalln(e)
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Bir sorun oluştu. :(")
}

func otherwiseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
	}

	wellcomeHandler(w, r)
}
func wellcomeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hoşgeldin :)")
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	written, err := json.Marshal(data.Pokemons)
	if err != nil {
		panic(err)
	}
	i, err := w.Write(written)
	if err != nil {
		fmt.Println("error after ")
		fmt.Println(i)
		panic(err)
	}
	log.Println("/list url:", r.URL)
	fmt.Fprint(w, "The List Handler\n")
}
