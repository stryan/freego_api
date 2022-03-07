package main

import (
	"errors"
	"fmt"

	"git.saintnet.tech/stryan/freego"
)

//Session represents an active game
type Session struct {
	simulator  *freego.Game
	redPlayer  *Player
	bluePlayer *Player
	moveNum    int
	moveList   []freego.ParsedCommand
}

//Player is a player in a match
type Player struct {
	Ready bool
	Team  freego.Colour
}

//ID returns player ID
func (p *Player) ID() int {
	panic("not implemented") // TODO: Implement
}

//Colour returns player team
func (p *Player) Colour() freego.Colour {
	return p.Team
}

//NewSession creates a new game session
func NewSession() *Session {
	sim := freego.NewGame()
	return &Session{
		simulator:  sim,
		redPlayer:  &Player{false, freego.Red},
		bluePlayer: &Player{false, freego.Blue},
		moveNum:    0,
		moveList:   []freego.ParsedCommand{},
	}
}

func (s *Session) tryMove(player *Player, move string) (*freego.ParsedCommand, error) {
	raw, err := freego.NewRawCommand(move)
	if err != nil {
		return nil, err
	}
	p, err := s.simulator.Parse(player, raw)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Session) mutate(p *freego.ParsedCommand) error {
	success, err := s.simulator.Mutate(p)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("invalid move")
	}
	s.moveList[s.moveNum] = *p
	s.moveNum++
	return nil
}

func (s *Session) getMove(p *Player, num int) (string, error) {
	if num < 0 || num > s.moveNum {
		return "", errors.New("invalid move number")
	}
	return fmt.Sprintf("%v %v", num, s.moveList[num].String()), nil
}
