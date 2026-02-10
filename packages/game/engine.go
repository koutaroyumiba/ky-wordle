// Package game : Wordle Game Package
package game

// NewGame : make that game
func NewGame(word string, maxTurns int) *State {
	return &State{
		Word:     word,
		MaxTurns: maxTurns,
	}
}
