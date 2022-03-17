import React from "react";
import { Cell } from "../types";
import styles from "../styles/board.module.css";

interface BoardProps {
  cells: Cell[];
  cellWidth: number;
}

const Board = (props: BoardProps) => {
  return (
    <div
      className={styles.gridContainer}
      style={{
        gridTemplateColumns: `repeat(${props.cellWidth}, 1fr)`,
      }}
    >
      {props.cells.map((cell) => (
        <div className={styles.gridCell}>{cell.piece}</div>
      ))}
    </div>
  );
};

export default Board;
