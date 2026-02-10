// Package game : Wordle Game Package
package game

// PlayerID : id of a player
type PlayerID string

// State : state of the game
type State struct {
	Word     string
	MaxTurns int
	Turn     int
	Finished bool
}
