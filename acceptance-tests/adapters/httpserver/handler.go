package httpserver

import (
	interactions "acceptance-tests/domain/interactions"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var interaction func(string) string

	switch r.URL.Path {
	case "/greet":
		interaction = interactions.Greet
	case "/curse":
		interaction = interactions.Curse
	default:
		http.NotFound(w, r)
	}

	fmt.Fprint(w, interaction(name))
}
