export interface Cell {
  piece?: string;
  terrain: boolean;
  hidden: boolean;
  empty: boolean;
}

export interface Position {
  row: number;
  column: number;
}