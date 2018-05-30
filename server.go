package pokedex

import (
	"log"
	"net/http"
	"strconv"
)

type Endpoint struct {
	Uri            string
	HandleFunction func(http.ResponseWriter, *http.Request)
}

// PokedexServer is a struct for representing server
type PokedexServer struct {
	Endpoints []*Endpoint
	Port      int
}

// Serve is for registering handlers
func (s *PokedexServer) Serve() {
	log.Println("starting server on :", s.Port)

	for _, e := range s.Endpoints {
		http.HandleFunc(e.Uri, e.HandleFunction)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}

// AddEndpoint adds new endpoint
func AddEndpoint(uri string, handler func(http.ResponseWriter, *http.Request)) {

}

// NewServer creates and returns a new server instance
func NewServer(port int) *PokedexServer {
	http.HandleFunc("/", otherwiseHandler)
	return &PokedexServer{Port: port}
}
