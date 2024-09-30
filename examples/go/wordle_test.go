package wordle

import (
	"errors"
	"fmt"
	"testing"
)

// Helper function to print color name in error logs
func getColor(color Color) string {
	switch color {
	case 0:
		return "Blank"
	case 1:
		return "Gray"
	case 2:
		return "Yellow"
	case 3:
		return "Green"
	default:
		return "Unknown"
	}
}

// Provide more informative string representation of Square struct for error logs
func (square Square) String() string {
	return fmt.Sprintf("Square{letter:%q, color:%s}", square.letter, getColor(square.color))
}

var testData = []struct {
	word     string
	guess    string
	expected []Square
}{
	{
		"apple",
		"apple",
		[]Square{
			{'a', Green},
			{'p', Green},
			{'p', Green},
			{'l', Green},
			{'e', Green},
		},
	},
	{
		"apple",
		"brink",
		[]Square{
			{'b', Gray},
			{'r', Gray},
			{'i', Gray},
			{'n', Gray},
			{'k', Gray},
		},
	},
	{
		"apple",
		"pears",
		[]Square{
			{'p', Yellow},
			{'e', Yellow},
			{'a', Yellow},
			{'r', Gray},
			{'s', Gray},
		},
	},
	{
		"stomp",
		"stamp",
		[]Square{
			{'s', Green},
			{'t', Green},
			{'a', Gray},
			{'m', Green},
			{'p', Green},
		},
	},
}

func TestInitialBoardIsAllBlank(t *testing.T) {
	t.Skip()
	game := NewGame("apple")
	blankSquare := Square{color: Blank}
	for i, row := range game.board {
		for j, square := range row {
			if square != blankSquare {
				t.Errorf("Expected blank square for row %d square %d,got %v", i, j, square)
			}
		}
	}
}

func TestCheckGuess(t *testing.T) {
	t.Skip()
	for _, td := range testData {
		game := NewGame(td.word)
		result, err := game.CheckGuess(td.guess)
		if err != nil {
			t.Errorf("Expected err to be nil, got: %s", err)
		}
		for i, resultSquare := range result {
			expectedSquare := td.expected[i]
			if resultSquare != expectedSquare {
				t.Errorf("Expected square %v does not match result square %v", expectedSquare, resultSquare)
			}
		}
	}
}

func TestCannotGuessSameWord(t *testing.T) {
	t.Skip()
	game := NewGame("apple")
	_, err := game.CheckGuess("pears")
	if err != nil {
		t.Errorf("Unexpected error on first guess: %v", err)
	}
	_, err = game.CheckGuess("pears")
	if !errors.Is(err, ReusedWord) {
		t.Errorf("Expected error `ReusedWord`; Acutal error: %v", err)
	}
}

func TestOnlySixGuesses(t *testing.T) {
	t.Skip()
	game := NewGame("apple")
	guesses := []string{"pears", "blank", "blink", "grape", "trust", "crate"}
	for _, guess := range guesses {
		_, err := game.CheckGuess(guess)
		if err != nil {
			t.Errorf("Unexpected error on guess %s: %v", guess, err)
		}
	}
	// Seventh guess should fail
	_, err := game.CheckGuess("crown")
	if !errors.Is(err, MaxAttempts) {
		t.Errorf("Expected error `MaxAttempts`; Actual error: %v", err)
	}
}

func TestCannotGuessAfterCorrectGuess(t *testing.T) {
	t.Skip()
	game := NewGame("apple")
	_, err := game.CheckGuess("apple")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	_, err = game.CheckGuess("blink")
	if !errors.Is(err, GameOver) {
		t.Errorf("Expected error: `GameOver`; Actual error: %v", err)
	}
}
