package main

import "math"

func encodeBoard(board [][]uint8) uint32 {
	var result uint32
	for i, row := range board {
		for j, v := range row {
			result = result + uint32(v)*uint32(math.Pow(3, float64(i*boardSize+j)))
		}
	}
	return result
}

func decodeBoard(code uint32) [][]uint8 {
	result := make([][]uint8, boardSize)
	for i := range result {
		result[i] = make([]uint8, boardSize)
	}

	for i, row := range result {
		for j := range row {
			result[i][j] = uint8((code / uint32(math.Pow(3, float64(i*boardSize+j)))) % 3)
		}
	}
	return result
}
