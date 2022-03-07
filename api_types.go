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
	GameBoard *freego.Game `json:"board"`
}

type gameStatusReq struct {
	PlayerID string `json:"player_id"`
}

type gameStatusResp struct {
	GameStatus freego.GameState `json:"game_status"`
	Move       int              `json:"move"`
}
