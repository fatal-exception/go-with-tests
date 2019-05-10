package main

import (
	"fmt"
	"net/http"
)

// InMemoryPlayerStore used for hardcoding player scores
type InMemoryPlayerStore struct{}

// GetPlayerScore does that
func (i *InMemoryPlayerStore) GetPlayerScore(name string) (res int, ok bool) {
	return 123, true
}

// PlayerServer is used to serve a player's scores
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)

}

// PlayerStore stores a player's scores
type PlayerStore interface {
	GetPlayerScore(name string) (res int, ok bool)
}
