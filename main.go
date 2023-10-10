package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, err := NewFileSystemPlayerStore(file)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
