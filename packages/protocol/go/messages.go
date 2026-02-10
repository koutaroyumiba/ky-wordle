// Package protocol : for all that protocol goodness
package protocol

// GuessRequest : request for a guess
type GuessRequest struct {
	Guess string `json:"guess"`
}
