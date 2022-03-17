import { Position } from "../types";

export const convertIndexToPosition = (
  index: number,
  cellWidth: number
): Position => ({
  row: Math.floor(index / cellWidth),
  column: index % cellWidth,
});
