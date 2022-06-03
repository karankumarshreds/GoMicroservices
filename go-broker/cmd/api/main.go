package main

import (
	"log"
	"net/http"
)

const PORT = ":8000"

type App struct{}

func main() {
	app := App{}
	log.Printf("Starting broker service on port %s\n", PORT)

	// define http server
	server := &http.Server{
		Addr:    PORT,
		Handler: app.routes(),
	}

	// start the server
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
