package main

import (
	"bufio"
	"log"
	"os"
)

var (
	X, M, A, S = byte(88), byte(77), byte(65), byte(83)
	mas        = []byte{M, A, S}
	xmasFound  = 0
	x_masFound = 0
)

func day4() {
	data := getDay4Data()

	for i := 0; i < len(data); i++ {
		len := len(data[i])
		for j := 0; j < len; j++ {
			row := data[i]
			if row[j] == X {
				checkXMAS(data, row, i, j)
			}
			if row[j] == A {
				checkX_MAS(data, i, j)
			}
		}
	}

	log.Printf(
		"XMAS Found: %d\nX_MAS Found: %d",
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

func checkXMAS(data []string, row string, i, j int) {
	checkLeft(row, j)
	checkRight(row, j)
	checkUp(data, i, j)
	checkDown(data, i, j)
	checkDownRight(data, i, j)
	checkDownLeft(data, i, j)
	checkUpLeft(data, i, j)
	checkUpRight(data, i, j)
}

func checkUpRight(data []string, i, j int) {
	if i-3 < 0 || j+4 > len(data[0]) {
		return
	}

	for k, char := range mas {
		if data[i-1-k][j+1+k] != char {
			return
		}
	}

	xmasFound++
}

func checkUpLeft(data []string, i, j int) {
	if i-3 < 0 || j-3 < 0 {
		return
	}

	for k, char := range mas {
		if data[i-1-k][j-1-k] != char {
			return
		}
	}

	xmasFound++
}

func checkDownLeft(data []string, i, j int) {
	if i+3 >= len(data) || j-3 < 0 {
		return
	}

	for k, char := range mas {
		if data[i+1+k][j-1-k] != char {
			return
		}
	}

	xmasFound++
}

func checkDownRight(data []string, i, j int) {
	if i+3 >= len(data) || j+4 > len(data[0]) {
		return
	}

	for k, char := range mas {
		if data[i+1+k][j+1+k] != char {
			return
		}
	}

	xmasFound++
}

func checkDown(data []string, i, j int) {
	if i+3 >= len(data) {
		return
	}

	for k, char := range mas {
		if data[i+1+k][j] != char {
			return
		}
	}

	xmasFound++
}

func checkUp(data []string, i, j int) {
	if i-3 < 0 {
		return
	}

	for k, char := range mas {
		if data[i-1-k][j] != char {
			return
		}
	}

	xmasFound++
}

func checkRight(row string, j int) {
	if j+4 > len(row) {
		return
	}
	for i, char := range mas {
		if row[j+i+1] != char {
			return
		}
	}

	xmasFound++
}

func checkLeft(row string, j int) {
	if j-3 < 0 {
		return
	}

	for i, char := range mas {
		if row[j-i-1] != char {
			return
		}
	}

	xmasFound++
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
