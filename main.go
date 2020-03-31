package main

import (
	"fmt"
	"math"
)

const boardSize = 4

type queueItem struct {
	boardCode uint32
	steps     uint32
}

func tryStep(i, j, dx, dy int, history []uint32, board [][]uint8) (canStep bool, newBoardCode uint32) {
	if i+dx >= 0 && i+dx < boardSize && j+dy >= 0 && j+dy < boardSize && board[i+dx][j+dy] != 0 {
		newBoard := make([][]uint8, boardSize)
		for k := range newBoard {
			newBoard[k] = make([]uint8, boardSize)
			copy(newBoard[k], board[k])
		}
		newBoard[i+dx][j+dy], newBoard[i][j] = newBoard[i][j], newBoard[i+dx][j+dy]
		newBoardCode = encodeBoard(newBoard)
		if history[newBoardCode] == 0 {
			canStep = true
		}
	}
	return
}

func displayResults(boardCode, initialBoardCode uint32, history []uint32) {
	if boardCode != initialBoardCode {
		displayResults(history[boardCode], initialBoardCode, history)
	}
	board := decodeBoard(boardCode)
	fmt.Println("=====")
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func main() {
	var initialBoardCode = encodeBoard([][]uint8{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{2, 2, 0, 2},
		{2, 2, 2, 2},
	})

	var desiredBoardCode = encodeBoard([][]uint8{
		{1, 1, 1, 1},
		{1, 1, 0, 1},
		{2, 0, 2, 2},
		{2, 2, 2, 2},
	})

	var maxCode = int(math.Pow(3, 16))

	var history = make([]uint32, maxCode+1)

	queue := make([]queueItem, 1)

	queue[0] = queueItem{boardCode: initialBoardCode, steps: 0}

	for len(queue) > 0 {
		dequeued := queue[0]
		queue = queue[1:]

		if dequeued.boardCode == desiredBoardCode {
			fmt.Printf("Solution found, %v steps\n", dequeued.steps)
			displayResults(desiredBoardCode, initialBoardCode, history)
			return
		}

		board := decodeBoard(dequeued.boardCode)

		for i, row := range board {
			for j := range row {
				if board[i][j] == 0 {
					for _, step := range [][]int{{-2, -1}, {-2, +1}, {-1, -2}, {-1, +2}, {+1, -2}, {+1, +2}, {+2, -1}, {+2, +1}} {
						dx, dy := step[0], step[1]
						if canStep, newBoardCode := tryStep(i, j, dx, dy, history, board); canStep {
							queue = append(queue, queueItem{boardCode: newBoardCode, steps: dequeued.steps + 1})
							history[newBoardCode] = dequeued.boardCode
						}
					}
				}
			}
		}
	}
	fmt.Println("No solution")

}
