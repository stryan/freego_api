package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"git.saintnet.tech/stryan/freego"
)

func respondWithError(res http.ResponseWriter, code int, message string) {
	respondWithJSON(res, code, map[string]string{"error": message})
}

func respondWithJSON(res http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	res.Write(response)
}

//TODO remove this when you can actually setup a game
func initDummy(g *freego.Game) {
	//Setup terrain
	terrain := []struct {
		x, y, t int
	}{
		{1, 1, 1},
		{2, 2, 1},
	}
	for _, tt := range terrain {
		res, err := g.Board.AddTerrain(tt.x, tt.y, tt.t)
		if err != nil {
			panic(err)
		}
		if !res {
			panic(errors.New("Error creating terrain"))
		}
	}
	pieces := []struct {
		x, y int
		p    *freego.Piece
	}{
		{0, 0, freego.NewPiece(freego.Flag, freego.Blue)},
		{3, 0, freego.NewPiece(freego.Spy, freego.Blue)},
		{2, 0, freego.NewPiece(freego.Captain, freego.Blue)},
		{3, 1, freego.NewPiece(freego.Marshal, freego.Blue)},
		{0, 1, freego.NewPiece(freego.Bomb, freego.Blue)},

		{1, 6, freego.NewPiece(freego.Flag, freego.Red)},
		{3, 6, freego.NewPiece(freego.Spy, freego.Red)},
		{2, 7, freego.NewPiece(freego.Captain, freego.Red)},
		{0, 6, freego.NewPiece(freego.Marshal, freego.Red)},
		{0, 7, freego.NewPiece(freego.Bomb, freego.Red)},
	}
	for _, tt := range pieces {
		res, err := g.SetupPiece(tt.x, tt.y, tt.p)
		if err != nil {
			panic(fmt.Errorf("Piece %v,%v:%v", tt.x, tt.y, err))
		}
		if !res {
			panic(errors.New("error placing dummy piece"))
		}
	}
	g.Start()
}
