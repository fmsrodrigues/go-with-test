package main

import (
	go_specs_greet "acceptance-tests"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	http.ListenAndServe(":8080", handler)
}
