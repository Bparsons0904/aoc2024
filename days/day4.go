package days

import (
	"bufio"
	"log"
	"os"
)

type Position struct {
	Row, Col int
}

var (
	X, M, A, S = byte(88), byte(77), byte(65), byte(83)
	mas        = []byte{M, A, S}
	xmasFound  = 0
	x_masFound = 0
	Directions = []Position{
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, 0},  // Up
		{1, 0},   // Down
		{1, 1},   // Down-Right
		{1, -1},  // Down-Left
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
	}
)

func Day4() {
	data := getDay4Data()

	for i := 0; i < len(data); i++ {
		len := len(data[i])
		for j := 0; j < len; j++ {
			row := data[i]
			if row[j] == X {
				checkXMAS(data, i, j)
			}
			if row[j] == A {
				checkX_MAS(data, i, j)
			}
		}
	}

	log.Printf(
		"XMAS Found: %d - X_MAS Found: %d",
		xmasFound,
		x_masFound,
	)
}

func checkX_MAS(data []string, i, j int) {
	if i-1 < 0 || j+2 > len(data) || i+2 > len(data) || j-1 < 0 {
		return
	}
	if ((data[i-1][j-1] == M && data[i+1][j+1] == S) || (data[i-1][j-1] == S && data[i+1][j+1] == M)) &&
		((data[i-1][j+1] == M && data[i+1][j-1] == S) || (data[i-1][j+1] == S && data[i+1][j-1] == M)) {
		x_masFound++
	}
}

func checkXMAS(data []string, rowIndex, colIndex int) {
	rowLen, colLen := len(data), len(data[0])
	for _, direction := range Directions {
		isValidPattern(data, rowIndex, colIndex, rowLen, colLen, direction)
	}
}

func isValidPattern(
	data []string,
	rowIndex int,
	colIndex int,
	rowLen int,
	colLen int,
	direction Position,
) {
	for i := range mas {
		newRow := rowIndex + direction.Row*(i+1)
		newCol := colIndex + direction.Col*(i+1)
		if !isInBounds(newRow, newCol, rowLen, colLen) {
			return
		}
	}

	for i, char := range mas {
		if data[rowIndex+direction.Row*(i+1)][colIndex+direction.Col*(i+1)] != char {
			return
		}
	}

	xmasFound++
}

func isInBounds(row, col, rowLen, colLen int) bool {
	return row >= 0 && row < rowLen && col >= 0 && col < colLen
}

func getDay4Data() []string {
	file, err := os.Open("inputs/input4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	results := []string{}
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results
}
