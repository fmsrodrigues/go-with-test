package main

import (
	poker "applications/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

// Use the following command to create a symlink to the html file
// for the /game endpoint, so that you can edit the html file and
// see the changes without having to copy the file to be served:
// $ ln -s ../../game.html game.html

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, err := poker.NewPlayerServer(store)
	if err != nil {
		log.Fatalf("could not create player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
