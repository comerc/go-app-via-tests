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
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
