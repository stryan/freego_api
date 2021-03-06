package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//API represents the api
type API struct {
	games   map[int]*Session
	nextInt int
}

//NewAPI creates new API instance
func NewAPI() *API {
	return &API{
		games:   make(map[int]*Session),
		nextInt: 1,
	}
}

//NewGame takes a POST and creates a new game or returns an open one
func (a *API) NewGame(res http.ResponseWriter, req *http.Request) {
	for i, g := range a.games {
		if !g.redPlayer.Ready {
			log.Println("red player somehow not ready")
			g.redPlayer.Ready = true
			if g.bluePlayer.Ready {
				g.simulator.Setup()
				initDummy(g.simulator)
				log.Println("dummy game started")
			}
			respondWithJSON(res, http.StatusOK, newGameResp{i, "red"})
			return
		}
		if !g.bluePlayer.Ready {
			g.bluePlayer.Ready = true
			if g.redPlayer.Ready {
				g.simulator.Setup()
				initDummy(g.simulator)
				log.Println("dummy game started")
			}
			respondWithJSON(res, http.StatusOK, newGameResp{i, "blue"})
			return
		}
	}
	log.Printf("creating new game %v", a.nextInt)
	a.games[a.nextInt] = NewSession(8)
	a.games[a.nextInt].redPlayer.Ready = true
	respondWithJSON(res, http.StatusOK, newGameResp{a.nextInt, "red"})
	a.nextInt = a.nextInt + 1
}

//GetGame returns current state of game, filtered accordingly
func (a *API) GetGame(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid game ID")
		return
	}
	pid := req.Header.Get("Player-id")
	var p *Player

	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	if pid == "red" {
		p = s.redPlayer
	} else if pid == "blue" {
		p = s.bluePlayer
	} else {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}

	log.Println("sending game state")
	board := s.getBoard(p)
	rotate := req.Header.Get("Rotate")
	if rotate == "true" {
		log.Println("rotating output")
		//rotateBoard(board)
		//rotateBoard(board)
		//rotateBoard(board)

	}
	respondWithJSON(res, http.StatusOK, gameResp{board})
	return
}

//GetGameStatus returns current game status and turn number
func (a *API) GetGameStatus(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid game ID")
		return
	}
	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	pid := req.Header.Get("Player-id")
	var p *Player
	if pid == "red" {
		p = s.redPlayer
	} else if pid == "blue" {
		p = s.bluePlayer
	} else {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	log.Printf("sending game status to player %v", p.Team)
	respondWithJSON(res, http.StatusOK, gameStatusResp{s.simulator.State, s.moveNum})
}

//PostMove attempts to make a game move
func (a *API) PostMove(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid game ID")
		return
	}
	var gr gameMovePostReq
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&gr); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer req.Body.Close()
	pid := req.Header.Get("Player-id")
	var p *Player

	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	if pid == "red" {
		p = s.redPlayer
	} else if pid == "blue" {
		p = s.bluePlayer
	} else {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	parsed, err := s.tryMove(p, gr.Move)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())

	}
	result, err := s.mutate(parsed)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
	}
	if result == "" {
		respondWithJSON(res, http.StatusOK, gameMovePostRes{true, false, "", err})
	}
	respondWithJSON(res, http.StatusOK, gameMovePostRes{true, true, result, err})
}

//GetMove returns the move made at turn X
func (a *API) GetMove(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid game ID")
		return
	}
	move, err := strconv.Atoi(vars["movenum"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid move number")
		return
	}
	var p *Player

	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	pid := req.Header.Get("Player-id")
	if pid == "red" {
		p = s.redPlayer
	} else if pid == "blue" {
		p = s.bluePlayer
	} else {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	moveRes, err := s.getMove(p, move)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "No such move")
		return
	}
	respondWithJSON(res, http.StatusOK, gameMoveRes{moveRes})
}
