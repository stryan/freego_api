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
