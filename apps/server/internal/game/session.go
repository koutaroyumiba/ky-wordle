package game

import (
	"sync"

	game "github.com/koutaroyumiba/ky-wordle/packages/game"
)

type Session struct {
	mu   sync.Mutex
	game *game.State
}

func NewSession(word string) *Session {
	return &Session{
		game: game.NewGame(word, 6),
	}
}

func (s *Session) Guess(player game.PlayerID, guess string) (game.GuessResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.game.SubmitGuess(player, guess)
}
