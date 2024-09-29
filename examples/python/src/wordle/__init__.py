from enum import Enum, auto


class ReusedWord(Exception):
    """Raise exception when player guesses word twice"""

    ...


class MaxAttempts(Exception):
    """Raise exception when player tries to make a 7th guess"""

    ...


class GameOver(Exception):
    """Raise exception when player tries to make guess after correctly guessing word"""


class Color(Enum):
    BLANK = auto()
    GRAY = auto()
    YELLOW = auto()
    GREEN = auto()


Square = tuple[str, Color]


class Wordle:
    def __init__(self, word: str) -> None:
        # TODO
        pass

    def check_guess(self, guess: str) -> list[Square]:
        # TODO
        return []
