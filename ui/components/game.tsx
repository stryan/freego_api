import React, { useState, useEffect } from "react";
import { fetchGameState, submitMove } from "../api/game.api";
import { convertIndexToPosition } from "../utils/position";
import { Cell } from "../types";
import Board from "./board";

const DEFAULT_CLICKED_CELL = -1;

interface GameProps {
  gameId: number;
  playerId: string;
}

const Game = (props: GameProps) => {
  const [isLoading, setLoading] = useState(false);
  const [cellWidth, setCellWidth] = useState(0);
  const [cells, setCells] = useState([] as Cell[]);
  const [focusedCellIndex, setFocusedCellIndex] =
    useState(DEFAULT_CLICKED_CELL);

  const onCellClicked = (cellIndex: number): void => {
    if (cellIndex === focusedCellIndex) {
      setFocusedCellIndex(DEFAULT_CLICKED_CELL);
      return;
    } else if (focusedCellIndex === DEFAULT_CLICKED_CELL) {
      setFocusedCellIndex(cellIndex);
      return;
    }
    submitMove(
      props.playerId,
      props.gameId,
      convertIndexToPosition(focusedCellIndex, cellWidth),
      convertIndexToPosition(cellIndex, cellWidth)
    );
  };

  useEffect(() => {
    setLoading(true);
    fetchGameState(props.playerId, props.gameId).then(
      ({ cells: cellList, cellWidth: width }) => {
        setCellWidth(width);
        setCells(cellList);
        setLoading(false);
      }
    );
  }, [props.gameId]);

  return (
    <Board
      cellWidth={cellWidth}
      cells={cells}
      focusedCellIndex={focusedCellIndex}
      onCellClick={onCellClicked}
    />
  );
};

export default Game;
