import React, { useState } from "react";
import cn from "classnames";
import { Cell } from "../types";
import styles from "../styles/board.module.css";

interface BoardProps {
  cells: Cell[];
  cellWidth: number;
  focusedCellIndex?: number;
  onCellClick: (index: number) => void;
}

const Board = (props: BoardProps) => (
  <div
    className={styles.gridContainer}
    style={{
      gridTemplateColumns: `repeat(${props.cellWidth}, 1fr)`,
    }}
  >
    {props.cells.map((cell, i) => (
      <button
        className={cn(styles.gridCell, {
          [styles.cellClicked]: i === props.focusedCellIndex,
        })}
        onClick={() => props.onCellClick(i)}
      >
        {cell.piece}
      </button>
    ))}
  </div>
);

export default Board;
