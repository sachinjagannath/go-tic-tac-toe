package main

import (
	"fmt"
	"log"

	"github.com/sachinjagannath/go-tic-tac-toe/game"
)

func main() {
	fmt.Println("\n==============================")
	fmt.Println(" TIC TAC TOE - Terminal Based")
	fmt.Println("==============================")
	game.WelcomeMessage()
	board := game.NewBoard()
	board.DisplayBoard()
	terminal := game.NewTerminal()
	g := game.AssignGame()
	if err := g.Play(terminal); err != nil {
		log.Fatal(err)
	}
}
