package wordle

import (
	"errors"
)

// Define new type `Color` as a type alias to `int`
type Color int

// Define Color enum values
const (
	Blank Color = iota
	Gray
	Yellow
	Green
)

// TODO
type Square struct {
	letter rune
	color  Color
}

// TODO
type Wordle struct{}

// TODO
func newGame(word string) Wordle {
	return Wordle{}
}

// TODO
func (wordle *Wordle) CheckGuess(guess string) ([]Square, error) {
	result := []Square{}
	return result, errors.New("Not yet implemented")
}
