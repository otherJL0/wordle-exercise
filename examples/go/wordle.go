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

var (
	Unimplemnted  = errors.New("This functionality has not been implemented")
	GuessTooShort = errors.New("Guess is shorter than 5 letters")
	GuessTooLong  = errors.New("Guess is longer than 5 letters")
	ReusedWord    = errors.New("Player already guessed this word")
	MaxAttempts   = errors.New("Player already guessed 6 times")
	GameOver      = errors.New("Player already guessed word, game over")
)

type Square struct {
	letter rune
	color  Color
}

type Wordle struct {
	word  string
	board [][]Square
}

// TODO
func NewGame(word string) Wordle {
	return Wordle{}
}

// TODO
func (wordle *Wordle) CheckGuess(guess string) ([]Square, error) {
	result := []Square{}
	return result, Unimplemnted
}
