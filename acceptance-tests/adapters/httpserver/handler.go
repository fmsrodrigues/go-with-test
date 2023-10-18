package httpserver

import (
	interactions "acceptance-tests/domain/interactions"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, interactions.Greet(name))
}
