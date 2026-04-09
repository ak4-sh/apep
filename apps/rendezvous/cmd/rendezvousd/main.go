package main

import (
	"log"
	"net/http"
	"time"

	rv "github.com/ak4-sh/apep/apps/rendezvous"
)

func main() {
	mux := http.NewServeMux()

	rv.RegisterRoutes(mux)

	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      http.NewCrossOriginProtection().Handler(mux),
	}

	log.Println("listening on port :8080...")
	log.Fatal(srv.ListenAndServe())
}
