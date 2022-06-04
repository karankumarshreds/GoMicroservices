package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `http:"error"`
	Message string      `http:"message"`
	Data    interface{} `http:"data,omitempty"`
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data *interface{}) error {
	maxBytes := 1000000 // 1mb : size we are allowed to read
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	err := json.NewDecoder(r.Body).Decode(&data)
	return err
}

func WriteJSON(w http.ResponseWriter, payload JsonResponse, statusCode int, headers ...http.Header) error {
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/http")
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Error while writing JSON", err)
	}
	return err
}

func ErrorJSON(w http.ResponseWriter, err error, statusCode ...int) {
	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()
	if len(statusCode) > 0 {
		WriteJSON(w, payload, statusCode[0])
	} else {
		WriteJSON(w, payload, http.StatusInternalServerError)
	}
}
