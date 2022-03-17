import wretch from "wretch";
import { Cell } from "../types";

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
