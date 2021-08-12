package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AntonioCorona-700/makeFood/pkg/api/router"
)

func main() {
	log.Println("Starting Food orientation")

	address := "localhost:8080"
	wt := 15 * time.Second
	rt := 15 * time.Second

	server := &http.Server{
		Handler:      router.Get(),
		Addr:         address,
		WriteTimeout: wt,
		ReadTimeout:  rt,
	}

	log.Printf("Listening at http://%s\n", address)
	log.Fatal(server.ListenAndServe())
}
