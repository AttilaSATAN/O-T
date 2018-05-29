package pokedex

import (
	"log"
	"net/http"
)

// PokedexServer is a struct for representing server
type PokedexServer struct {
}

// Serve is for registering handlers
func (s *PokedexServer) Serve() {
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)

}
