import wretch from "wretch";
import { Cell, Position } from "../types";

const USER_ID_HEADER = "Player-id";

export const fetchGameState = async (
  playerId: string,
  gameId: number
): Promise<{ cells: Cell[]; cellWidth: number }> => {
  const response: { board: Cell[][] } = await wretch(`/api/game/${gameId}`)
    .headers({
      [USER_ID_HEADER]: playerId,
    })
    .get()
    .json();
  return {
    cells: response.board.flat(),
    cellWidth: response.board[0].length,
  };
};

export const submitMove = async (
  playerId: string,
  gameId: number,
  piecePosition: Position,
  movePosition: Position
): Promise<{ cells: Cell[] }> => {
  console.log(piecePosition, movePosition);
  const response: { board: Cell[][] } = await wretch(`/api/game/${gameId}/move`)
    .headers({
      [USER_ID_HEADER]: playerId,
    })
    .body({
      pieceRow: piecePosition.row,
      pieceColumn: piecePosition.column,
      moveRow: movePosition.row,
      moveColumn: movePosition.column,
    })
    .post()
    .json();
  return {
    cells: response.board.flat(),
  };
};
