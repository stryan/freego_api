package main

//ViewTile is a json friendly version of a tile
type ViewTile struct {
	Piece   string `json:"piece"`
	Terrain bool   `json:"terrain"`
	Hidden  bool   `json:"hidden"`
	Empty   bool   `json:"empty"`
}

//NewViewTile creates a new ViewTile
func NewViewTile() *ViewTile {
	return &ViewTile{}
}
