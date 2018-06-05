package pokedex

import (
	"log"
	"net/http"
	"strconv"
)

//Endpoint is a represantation for an API end-point. An end-point is a single URI
type Endpoint struct {
	URI            string
	HandleFunction func(http.ResponseWriter, *http.Request)
}

// Server is a struct for representing server
type Server struct {
	Endpoints []Endpoint `json:"uri`
	Port      int
}

// Serve is initiation point for server. After defining the endpoints with AddEndpoint call this function.
func (s *Server) Serve() {

	log.Println("starting server on :", s.Port)

	for _, endpoint := range s.Endpoints {
		http.HandleFunc(endpoint.URI, endpoint.HandleFunction)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}

// AddEndpoint adds new endpoint
func (s *Server) AddEndpoint(URI string, handler func(http.ResponseWriter, *http.Request)) {
	s.Endpoints = append(s.Endpoints, Endpoint{URI, handler})
}

// NewServer creates and returns a new server instance
func NewServer(port int) *Server {

	return &Server{Port: port}
}
