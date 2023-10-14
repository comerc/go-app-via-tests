package main

import (
	"fmt"
	"log"
	"os"
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
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
