package main

import (
	"errors"
	"fmt"
	"log"

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
		moveNum:    1,
		moveList:   make([]freego.ParsedCommand, 20),
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

func (s *Session) mutate(p *freego.ParsedCommand) (string, error) {
	success, err := s.simulator.Mutate(p)
	if err != nil {
		return "", err
	}
	if success {
		s.moveList[s.moveNum] = *p
		s.moveNum++
		return fmt.Sprintf("%v %v", s.moveNum-1, p.String()), nil
	}
	return "", nil
}

func (s *Session) getMove(p *Player, num int) (string, error) {
	if num <= 0 || num >= s.moveNum {
		log.Printf("tried to get move number %v when move is %v", num, s.moveNum)
		return "", errors.New("invalid move number")
	}
	return fmt.Sprintf("%v %v", num, s.moveList[num].String()), nil
}

func (s *Session) getBoard(p *Player) [8][8]*ViewTile {
	var res [8][8]*ViewTile
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cur := NewViewTile()
			terrain, err := s.simulator.Board.IsTerrain(i, j)
			if err != nil {
				panic(err)
			}
			if terrain {
				cur.Terrain = true
			} else {
				piece, err := s.simulator.Board.GetPiece(i, j)
				if err != nil {
					panic(err)
				}
				if piece != nil {
					if piece.Hidden {
						cur.Hidden = true
						if piece.Owner == p.Colour() {
							cur.Piece = piece.Rank.String()
						} else {
							cur.Piece = "Unknown"
						}
					} else {
						cur.Piece = piece.Rank.String()
					}
				} else {
					cur.Empty = true
				}
			}
			res[i][j] = cur
		}
	}
	return res
}
