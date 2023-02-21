// func ListenAndServe(addr string, handler Handler) error - start a web server listening on a port, creating a goroutine for every request and running it against a Handler.
//If there is a problem the web server will return an error, an example of that might be the port already being listened to. For that reason we wrap the call in log.Fatal to log the error to the user.

//	type Handler interface {
//		ServeHTTP(ResponseWriter, *Request)
//	}
//
// A type implements the Handler interface in order to make a server, by implementing the ServeHTTP method which expects two arguments,
// the first is where we write our response and the second is the HTTP request that was sent to the server.
// We can do that by creating a struct and make it implement the interface by implementing its own ServeHTTP method.

// the use-case for structs is for holding data but currently we have no state, so it doesn't feel right to be creating one. Instead, we can use HandlerFunc
// type HandlerFunc has already implemented the ServeHTTP method. By type casting our PlayerServer function with it, we have now implemented the required Handler.
// frmo the docs - The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
package main

import (
	"fmt"
	"net/http"
	"strings"
)

// a funtion implementing the functionality of ServeHTTP that can be wrapped in HadlerFunc
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")

// 	fmt.Fprint(w, GetPlayerScore(player)) // writing response to http.ResponseWriter using fmt.Fprint
// }

type PlayerStore interface {
	GetPlayerScore(name string) int // we have not implemeted these but just calling them
	RecordWin(name string)
}

// PlayerServer can be a class (API i.e controller layer) - that calls methods GetPlayerScore(), IncrementPlayerScore() (service layer)
// Also, to separate the concerns for each player, GetPlayerScore(), IncrementPlayerScore() methods of a player can be grouper together using an interface PlayerStore
// So PlayerServer instead of referencing GetPlayerScore() directly,  will access it and other such player methods via PlayerStore

type PlayerServer struct {
	store PlayerStore // needs to use PlayerStore so reference present here
}

// implement ServeHTTP as hadler function - add a method to PlayerServer class/ struct
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound) // default is http.StatusOK
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// we can separate concerns for each player using an interface, PlayerStore
// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}

// 	return ""
// }

// func RecordWin(name string) {

// }
