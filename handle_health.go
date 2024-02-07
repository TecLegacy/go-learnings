package main

import "net/http"

func handleHealth(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
