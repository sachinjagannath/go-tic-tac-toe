package game

import (
	"bufio"
	"os"
)

type Terminal struct {
	reader *bufio.Reader
}

func NewTerminal() *Terminal {
	return &Terminal{
		reader: bufio.NewReader(os.Stdin),
	}
}
