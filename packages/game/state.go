package game

type PlayerID string

type State struct {
	Word     string
	MaxTurns int
	Turn     int
	Finished bool
	Winner   *PlayerID
}
