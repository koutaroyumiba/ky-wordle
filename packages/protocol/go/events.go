package protocol

import "github.com/koutaroyumiba/ky-wordle/packages/game"

type GuessResultEvent struct {
	Result game.GuessResult `json:"result"`
}

type GameOverEvent struct {
	Winner string `json:"winner"`
}
