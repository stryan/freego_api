import React, { useState, useEffect } from "react";
import { fetchGameState } from "../api/game.api";
import { Cell } from "../types";
import Board from "./board";

interface GameProps {
  gameId: number;
}

const Game = (props: GameProps) => {
  const [isLoading, setLoading] = useState(false);
  const [cellWidth, setCellWidth] = useState(4);
  const [cells, setCells] = useState([] as Cell[]);
  useEffect(() => {
    setLoading(true);
    fetchGameState("red", props.gameId).then(
      ({ cells: cellList, cellWidth: width }) => {
        setCellWidth(width);
        setCells(cellList);
        setLoading(false);
      }
    );
  }, [props.gameId]);

  return <Board cellWidth={cellWidth} cells={cells} />;
};

export default Game;
