package game

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	cells [3][3]rune
}

type Game struct {
	board         *Board
	player1       Player
	player2       Player
	currentPlayer Player
	moveCount     int
}

func NewBoard() *Board {
	b := &Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.cells[i][j] = '_'
		}
	}
	return b
}

func (b *Board) DisplayBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", b.cells[i][j])
		}
		fmt.Println()
	}
}

func AssignGame() *Game {
	player1 := NewPlayer("Player 1", 'X')
	player2 := NewPlayer("Player 2", 'O')

	return &Game{
		board:         NewBoard(),
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
		moveCount:     0,
	}
}

func (g *Game) Play(t *Terminal) error {
	fmt.Println("Enter the digits from 1 to 9...")

	for {
		if g.board.isFull() {
			fmt.Println("The match is drawn...")
			os.Exit(0)
		}

		input, err := t.reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Input should be a number...")
			continue
		}

		if num < 1 || num > 9 {
			fmt.Println("Input should be between 1 and 9...")
			continue
		}

		row := (num - 1) / 3
		col := (num - 1) % 3

		if !g.board.isEmpty(row, col) {
			fmt.Println("The cell is already taken! Try again...")
			continue
		}

		g.board.cells[row][col] = g.currentPlayer.Symbol
		g.board.DisplayBoard()
		if g.winner() {
			fmt.Printf("The winner is %s", g.currentPlayer.Name)
			os.Exit(0)
		}
		g.switchPlayer()
	}
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.player1 {
		fmt.Println("Player 2:")
		g.currentPlayer = g.player2
	} else {
		fmt.Println("Player 1:")
		g.currentPlayer = g.player1
	}
}

func (g *Game) winner() bool {
	symbol := g.currentPlayer.Symbol
	for i := 0; i < 3; i++ {
		if g.board.cells[i][0] == symbol &&
			g.board.cells[i][1] == symbol &&
			g.board.cells[i][2] == symbol {
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if g.board.cells[0][i] == symbol &&
			g.board.cells[1][i] == symbol &&
			g.board.cells[2][i] == symbol {
			return true
		}
	}

	if g.board.cells[0][0] == symbol &&
		g.board.cells[1][1] == symbol &&
		g.board.cells[2][2] == symbol {
		return true
	}

	if g.board.cells[0][2] == symbol &&
		g.board.cells[1][1] == symbol &&
		g.board.cells[2][0] == symbol {
		return true
	}

	return false
}

func WelcomeMessage() {
	fmt.Println()
	fmt.Println("Rules:")
	fmt.Println("-> Players take turns marking cells")
	fmt.Println("-> First to get 3 in a row wins")
	fmt.Println("-> Row, columns or diagonal counts")
	fmt.Println()
}

func (b *Board) isEmpty(row, col int) bool {
	return b.cells[row][col] == '_'
}

func (b *Board) isFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == '_' {
				return false
			}
		}
	}
	return true
}
