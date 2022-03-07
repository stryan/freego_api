package main

import "git.saintnet.tech/stryan/freego"

//Session represents an active game
type Session struct {
	simulator *freego.Game
	redReady  bool
	blueReady bool
	moveNum   int
}

//NewSession creates a new game session
func NewSession() *Session {
	return &Session{
		simulator: freego.NewGame(),
		redReady:  false,
		blueReady: false,
		moveNum:   0,
	}
}
