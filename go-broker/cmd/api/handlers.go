package main

import (
	"log"
	"net/http"
)

func (a *App) BrokerHandler(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	err := a.WriteJSON(w, payload, http.StatusOK)
	if err != nil {
		log.Println("Unable to send response to the client", err)
	}
}
