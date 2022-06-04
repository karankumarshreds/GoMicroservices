package main

import "net/http"

func (a *App) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := r.Read
}
