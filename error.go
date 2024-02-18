package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Printf("Error code 5XX %v", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, statusCode, errorResponse{
		Error: msg,
	})
}
