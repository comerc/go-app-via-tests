package main

import (
	"log"
	"net/http"
	"path/filepath"

	poker "github.com/comerc/go-app-via-tests"
)

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(
		filepath.Join("../..", poker.DBFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
