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
	Endpoints []Endpoint
	Port      int
}

type Meta struct {
	Total int
}

type Response struct {
	Data []Pokemon
	Meta Meta
}

// Serve is for registering handlers
func (s *PokedexServer) Serve() {

	log.Println("starting server on :", s.Port)

	for _, endpoint := range s.Endpoints {
		http.HandleFunc(endpoint.Uri, endpoint.HandleFunction)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}

// AddEndpoint adds new endpoint
func (s *PokedexServer) AddEndpoint(uri string, handler func(http.ResponseWriter, *http.Request)) {
	s.Endpoints = append(s.Endpoints, Endpoint{uri, handler})
}

// NewServer creates and returns a new server instance
func NewServer(port int) *PokedexServer {

	return &PokedexServer{Port: port}
}
