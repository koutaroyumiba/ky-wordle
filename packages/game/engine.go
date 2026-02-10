package game

import "errors"

var (
	ErrGameOver     = errors.New("game already finished")
	ErrInvalidGuess = errors.New("invalid guess length")
)

type LetterResult int

const (
	Absent LetterResult = iota
	Present
	Correct
)

type GuessResult struct {
	Letters []LetterResult
}

func NewGame(word string, maxTurns int) *State {
	return &State{
		Word:     word,
		MaxTurns: maxTurns,
	}
}

func (g *State) SubmitGuess(player PlayerID, guess string) (GuessResult, error) {
	if g.Finished {
		return GuessResult{}, ErrGameOver
	}

	if len(guess) != len(g.Word) {
		return GuessResult{}, ErrInvalidGuess
	}

	result := evaluateGuess(g.Word, guess)

	if isCorrectGuess(result) {
		g.Finished = true
		g.Winner = &player
	}

	g.Turn++
	return result, nil
}

func evaluateGuess(word, guess string) GuessResult {
	result := make([]LetterResult, len(word))
	counts := map[byte]int{}

	// find all green first
	for i := range guess {
		if word[i] == guess[i] {
			result[i] = Correct
		} else {
			counts[guess[i]]++
		}
	}

	// second pass for yellow
	for i := range guess {
		if result[i] == Correct {
			continue
		}

		if counts[guess[i]] > 0 {
			result[i] = Present
			counts[guess[i]]--
		} else {
			result[i] = Absent
		}
	}

	return GuessResult{Letters: result}
}

func isCorrectGuess(r GuessResult) bool {
	for _, l := range r.Letters {
		if l != Correct {
			return false
		}
	}

	return true
}
