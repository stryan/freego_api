package main

import "git.saintnet.tech/stryan/freego"

//type newGameReq struct{}

type newGameResp struct {
	GameID   int    `json:"game_id"`
	PlayerID string `json:"player_id"`
}

type gameReq struct {
	PlayerID string `json:"player_id"`
}

type gameResp struct {
	GameBoard [8][8]*ViewTile `json:"board"`
}

type gameStatusReq struct {
	PlayerID string `json:"player_id"`
}

type gameStatusResp struct {
	GameStatus freego.GameState `json:"game_status"`
	Move       int              `json:"move"`
}

type gameMovePostReq struct {
	PlayerID string `json:"player_id"`
	Move     string `json:"move"`
}

type gameMovePostRes struct {
	Valid  bool   `json:"valid"`
	Result bool   `json:"result"`
	Parsed string `json:"parsed"`
	Error  error  `json:"error"`
}

type gameMoveReq struct {
	PlayerID string `json:"player_id"`
}

type gameMoveRes struct {
	Move string `json:"move"`
}
