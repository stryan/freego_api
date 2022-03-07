package main

import (
	"encoding/json"
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

//NewGame takes a POST and creates a new game
func (a *API) NewGame(res http.ResponseWriter, req *http.Request) {
	a.games[a.nextInt] = NewSession()
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
	var gr gameReq
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&gr); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer req.Body.Close()
	if gr.PlayerID != "red" && gr.PlayerID != "blue" {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	//TODO filter based off player info
	respondWithJSON(res, http.StatusOK, gameResp{s.simulator})
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
	var gr gameReq
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&gr); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer req.Body.Close()
	if gr.PlayerID != "red" && gr.PlayerID != "blue" {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
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
		respondWithError(res, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer req.Body.Close()
	if gr.PlayerID != "red" && gr.PlayerID != "blue" {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	var p *Player

	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	if gr.PlayerID == "red" {
		p = s.redPlayer
	} else {
		p = s.bluePlayer
	}
	parsed, err := s.tryMove(p, gr.Move)
	if err != nil {
		respondWithJSON(res, http.StatusOK, gameMovePostRes{false, false, parsed.String(), err})
	}

}

//GetMove returns the move made at turn X
func (a *API) GetMove(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid game ID")
		return
	}
	move, err := strconv.Atoi(vars["move_number"])
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid move number")
		return
	}
	var gr gameMoveReq
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&gr); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer req.Body.Close()
	if gr.PlayerID != "red" && gr.PlayerID != "blue" {
		respondWithError(res, http.StatusBadRequest, "Bad player ID")
		return
	}
	var p *Player

	s, isset := a.games[id]
	if !isset {
		respondWithError(res, http.StatusBadRequest, "No such game")
		return
	}
	if gr.PlayerID == "red" {
		p = s.redPlayer
	} else {
		p = s.bluePlayer
	}
	moveRes, err := s.getMove(p, move)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, "No such move")
		return
	}
	respondWithJSON(res, http.StatusOK, gameMoveRes{moveRes})
}
