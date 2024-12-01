package _0_http

import "io"

// Game manages the state of a game.
type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}
