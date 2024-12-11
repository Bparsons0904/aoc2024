package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
)

type Trail struct {
	Data            [][]int
	CurrentPosition Position
	PreviousPath    []Position
}

func Day10() {
	data := getDay10Data()
	trailHeads := getTrailHeads(data)
	fullTrailHeads := processTrailheads(data, trailHeads)
	log.Println(fullTrailHeads)
}

func getTrailHeads(data [][]int) []Position {
	positions := []Position{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[i][j] == 0 {
				positions = append(positions, Position{Row: i, Col: j})
			}
		}
	}

	return positions
}

func processTrailheads(data [][]int, trailHeads []Position) int {
	dataPosition := Position{len(data), len(data[0])}

	results := 0
	// log.Println(trailHeads)
	for _, trailHead := range trailHeads {
		// if i == 0 {
		// path := []Position{
		// 	trailHead,
		// }
		result := len(checkPaths(trailHead, dataPosition, data, []Position{}))
		// log.Println("results from ", trailHead, result)
		results += result
		// }
	}
	return results
}

func checkPaths(
	starting Position,
	dataPosition Position,
	data [][]int,
	completedPaths []Position,
) []Position {
	pathsInBounds := []Position{}
	for _, direction := range directions[:4] {
		newPosition := starting.MoveTo(direction)
		if checkPositionInBounds(newPosition, dataPosition) {
			pathsInBounds = append(pathsInBounds, newPosition)
		}

	}
	potentialMoves := []Position{}
	for _, pathToCheck := range pathsInBounds {
		if data[starting.Row][starting.Col] == 8 &&
			data[pathToCheck.Row][pathToCheck.Col] == 9 {
			completedPaths = append(completedPaths, pathToCheck)
		}
		if data[starting.Row][starting.Col]-data[pathToCheck.Row][pathToCheck.Col] == -1 {
			potentialMoves = append(potentialMoves, pathToCheck)
		}
	}

	if len(potentialMoves) == 0 {
		return completedPaths
	}

	results := []Position{}
	for _, move := range potentialMoves {
		foundPaths := checkPaths(move, dataPosition, data, completedPaths)
		for _, foundPath := range foundPaths {
			if !foundPath.Contains(results) {
				results = append(results, foundPath)
			}
		}
	}

	return results
}

func (position Position) Contains(positions []Position) bool {
	return slices.ContainsFunc(positions, func(test Position) bool {
		return test.Row == position.Row && test.Col == position.Col
	})
}

func (position Position) Equals(secondaryPosition Position) bool {
	return position.Row == secondaryPosition.Row && position.Col == secondaryPosition.Col
}

func (position Position) MoveTo(direction Position) Position {
	newRow, newCol := position.Row+direction.Row, position.Col+direction.Col
	return Position{
		newRow, newCol,
	}
}

func checkPositionInBounds(position Position, mapSize Position) bool {
	return isInBounds(position.Row, position.Col, mapSize.Row, mapSize.Col)
}

func getDay10Data() [][]int {
	file, err := os.Open("inputs/input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		intRow := []int{}

		for _, val := range row {
			num, err := strconv.Atoi(string(val))
			if err != nil {
				log.Fatal(err)
			}
			intRow = append(intRow, num)
		}

		results = append(results, intRow)
	}

	return results
}
