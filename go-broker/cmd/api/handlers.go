package main

import (
	"log"
	"net/http"
	common "github.com/karankumarshreds/GoMicroservices/go-common/http"
)

func (a *App) BrokerHandler(w http.ResponseWriter, r *http.Request) {
	payload := common.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	err := common.WriteJSON(w, payload, http.StatusOK)
	if err != nil {
		log.Println("Unable to send response to the client", err)
	}
}
