package main

import "net/http"

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
