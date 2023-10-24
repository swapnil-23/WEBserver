package main

import (
	"log"
	"net/http"
	server "server/pkg/SERVER"
	"time"
)

func main() {
	address := ":8080"
	mux := http.NewServeMux()
	srv := server.New()

	mux.HandleFunc("/", srv.HandleIndex)

	mux.HandleFunc("/users", srv.HandleUser)

	s := http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Start server at port 8080: %v", address)
	log.Fatal(s.ListenAndServe())
}
