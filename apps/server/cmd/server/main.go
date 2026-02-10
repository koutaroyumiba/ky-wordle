// Package main : main entrypoint for ky-wordle
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/koutaroyumiba/ky-wordle/apps/server/internal/game"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	session := game.NewSession("apple")
	http.HandleFunc("/guess", func(w http.ResponseWriter, r *http.Request) {
		guess := r.URL.Query().Get("g")
		result, err := session.Guess("p1", guess)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Fprintf(w, "%v\n", result)
	})

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
