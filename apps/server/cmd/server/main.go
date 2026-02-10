// Package main : main entrypoint for ky-wordle
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
