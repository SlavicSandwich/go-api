package main

import (
	"fmt"
	poker "go-api"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
