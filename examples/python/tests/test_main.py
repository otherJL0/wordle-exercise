from wordle import Color, Wordle, Square, ReusedWord, MaxAttempts, GameOver
import pytest

testdata = [
    (
        "apple",
        "apple",
        [
            ("a", Color.GREEN),
            ("p", Color.GREEN),
            ("p", Color.GREEN),
            ("l", Color.GREEN),
            ("e", Color.GREEN),
        ],
    ),
    (
        "apple",
        "brink",
        [
            ("b", Color.GRAY),
            ("r", Color.GRAY),
            ("i", Color.GRAY),
            ("n", Color.GRAY),
            ("k", Color.GRAY),
        ],
    ),
    (
        "apple",
        "pears",
        [
            ("p", Color.YELLOW),
            ("e", Color.YELLOW),
            ("a", Color.YELLOW),
            ("r", Color.GRAY),
            ("s", Color.GRAY),
        ],
    ),
    (
        "stomp",
        "stamp",
        [
            ("s", Color.GREEN),
            ("t", Color.GREEN),
            ("a", Color.GRAY),
            ("m", Color.GREEN),
            ("p", Color.GREEN),
        ],
    ),
]


@pytest.mark.skip()
def test_initial_board_is_all_blank():
    game = Wordle("apple")
    blank_row = [("", Color.BLANK)] * 5
    for row in game.board:
        assert row == blank_row


@pytest.mark.skip()
@pytest.mark.parametrize("word,guess,expected", testdata)
def test_guess(word: str, guess: str, expected: list[Square]):
    game = Wordle(word)
    result = game.check_guess(guess)
    assert result == expected


@pytest.mark.skip()
def test_cannot_guess_same_word():
    game = Wordle("apple")
    _ = game.check_guess("pears")
    with pytest.raises(ReusedWord):
        game.check_guess("pears")


@pytest.mark.skip()
def test_only_six_guesses():
    game = Wordle("apple")
    for guess in ["pears", "blank", "blink", "grape", "trust", "crate"]:
        _ = game.check_guess(guess)
    with pytest.raises(MaxAttempts):
        _ = game.check_guess("crown")


@pytest.mark.skip()
def test_cannot_guess_after_correctly_guessing_word():
    game = Wordle("apple")
    _ = game.check_guess("apple")
    with pytest.raises(GameOver):
        _ = game.check_guess("blink")
