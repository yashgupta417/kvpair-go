package main

import (
	"io"
	"log"
	"net/http"

	"github.com/yashgupta417/kvpair-go/handlers"
)

var server *http.Server

func main() {
	rootHandler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		io.WriteString(w, "Welcome to KV DB.")
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/key", handlers.GetKey)
	http.HandleFunc("/key/set", handlers.UpsertKey)
	http.HandleFunc("/key/remove", handlers.DeleteKey)

	log.Printf("Starting server....")

	server = &http.Server{
		Addr:    ":8001",
		Handler: nil,
	}
	server.ListenAndServe()
}