server:
	cd ./apps/server && go run ./cmd/server

terminal:
	cd ./apps/tui && go run ./cmd/tui

.PHONY: server terminal
