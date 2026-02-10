module github.com/koutaroyumiba/ky-wordle/apps/server

go 1.25.6

replace github.com/koutaroyumiba/ky-wordle/packages/game => ../../packages/game

replace github.com/koutaroyumiba/ky-wordle/packages/protocol => ../../packages/protocol

require github.com/koutaroyumiba/ky-wordle/packages/game v0.0.0-00010101000000-000000000000
