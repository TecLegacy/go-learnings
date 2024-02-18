package main

import "net/http"

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
