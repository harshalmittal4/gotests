package main

import (
	"log"
	"net/http"
)

// see in_memory_player_store.go
// InMemoryPlayerStore collects data about players in memory.
// type InMemoryPlayerStore struct{}

// Retrieves score for a given player
// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return 123 // returns 123 for all players
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server)) // handler argument needs to be a typeo that has a ServeHTTP method, PlayerServer has that so can be used as hadler
}
