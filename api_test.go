package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestNewAPI(t *testing.T) {
	a := NewAPI()
	if len(a.games) != 0 {
		t.Fatalf("games list not empty")
	}
	if a.nextInt != 1 {
		t.Fatalf("nextInt somehow already implemented")
	}
}

func dummyGame(a *API) int {
	i := a.nextInt
	a.games[i] = NewSession(8)
	a.games[i].redPlayer.Ready = true
	a.games[i].bluePlayer.Ready = true
	a.games[i].simulator.Setup()
	initDummy(a.games[i].simulator)
	return i
}

func TestNewGame(t *testing.T) {
	t.Parallel()
	a := NewAPI()
	var tests = []struct {
		gid int
		pid string
	}{
		{1, "red"},
		{1, "blue"},
		{2, "red"},
	}
	for i, tt := range tests {
		tname := fmt.Sprintf("/game %v", i)
		t.Run(tname, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/game", nil)
			w := httptest.NewRecorder()
			a.NewGame(w, r)
			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Fatal("failed to create new game")
			}
			_, ok := a.games[tt.gid]
			if !ok {
				t.Fatalf("API thinks it created a game but it didn't")
			}
			var respStruct newGameResp
			err := json.NewDecoder(resp.Body).Decode(&respStruct)
			if err != nil {
				t.Errorf("/game returned bad response body: %v", err)
			}
			if respStruct.GameID != tt.gid {
				t.Errorf("Expected game %v, got %v", tt.gid, respStruct.GameID)
			}
			if respStruct.PlayerID != tt.pid {
				t.Errorf("wrong playerID returned")
			}
		})
	}
}

func TestGetGame(t *testing.T) {
	t.Parallel()
	a := NewAPI()
	gid := dummyGame(a)
	var tests = []struct {
		pid  string
		code int
	}{
		{"red", http.StatusOK},
		{"blue", http.StatusOK},
		{"green", http.StatusBadRequest},
	}
	for _, tt := range tests {
		tname := fmt.Sprintf("/game from player %v", tt.pid)
		t.Run(tname, func(t *testing.T) {
			r, _ := http.NewRequest("GET", "/game", nil)
			w := httptest.NewRecorder()
			r = mux.SetURLVars(r, map[string]string{"id": strconv.Itoa(gid)})
			r.Header.Add("Player-id", tt.pid)
			r.Header.Add("Rotate", "false")
			a.GetGame(w, r)
			resp := w.Result()
			if resp.StatusCode != tt.code {
				t.Fatalf("failed to get game: %v", resp.Status)
			}
			if resp.StatusCode == http.StatusOK {
				var respStruct gameResp
				err := json.NewDecoder(resp.Body).Decode(&respStruct)
				if err != nil {
					t.Fatalf("/game returned bad response body: %v", err)
				}
				if len(respStruct.GameBoard) == 0 {
					t.Errorf("bad game board returned")
				}
				for j := range respStruct.GameBoard {
					for i, vt := range respStruct.GameBoard[j] {
						curr, err := a.games[gid].simulator.Board.GetPiece(i, j)
						if err != nil {
							t.Fatalf("Strange board position: %v", err)
						}
						if curr != nil && !vt.Hidden && curr.Owner.String() != tt.pid && curr.Hidden {
							t.Errorf("/game returned a piece that should be hidden but isn't")
						}
					}
				}
			}
		})
	}
}
