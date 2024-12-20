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
	fullTrailHeads, expandedTrailHeads := processTrailheads(data, trailHeads)
	log.Printf(
		"Found %d optimal hiking trails. Found %d optimal hiking paths.",
		fullTrailHeads,
		expandedTrailHeads,
	)
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

func processTrailheads(data [][]int, trailHeads []Position) (int, int) {
	dataPosition := Position{len(data), len(data[0])}

	results, expandedResults := 0, 0
	for _, trailHead := range trailHeads {
		results += len(checkPaths(trailHead, dataPosition, data, []Position{}))
		expandedResults += len(
			checkExpandedPath(trailHead, dataPosition, data, []Position{trailHead}, [][]Position{}),
		)
	}
	return results, expandedResults
}

func checkExpandedPath(
	starting Position,
	dataPosition Position,
	data [][]int,
	currentPath []Position,
	completedPaths [][]Position,
) [][]Position {
	newPath := append([]Position{}, currentPath...)
	newPath = append(newPath, starting)

	pathsInBounds := []Position{}
	for _, direction := range Directions[:4] {
		newPosition := starting.MoveTo(direction)
		if checkPositionInBounds(newPosition, dataPosition) {
			pathsInBounds = append(pathsInBounds, newPosition)
		}
	}

	potentialMoves := []Position{}
	for _, pathToCheck := range pathsInBounds {
		if data[starting.Row][starting.Col] == 8 &&
			data[pathToCheck.Row][pathToCheck.Col] == 9 {
			newCompletedPath := append([]Position{}, newPath...)
			newCompletedPath = append(newCompletedPath, pathToCheck)
			completedPaths = append(completedPaths, newCompletedPath)
		}
		if data[starting.Row][starting.Col]-data[pathToCheck.Row][pathToCheck.Col] == -1 {
			if !pathToCheck.IncludedIn(newPath) {
				potentialMoves = append(potentialMoves, pathToCheck)
			}
		}
	}

	if len(potentialMoves) == 0 {
		return completedPaths
	}

	results := completedPaths
	for _, move := range potentialMoves {
		foundPaths := checkExpandedPath(move, dataPosition, data, newPath, results)
		for _, foundPath := range foundPaths {
			if !containsPath(results, foundPath) {
				results = append(results, foundPath)
			}
		}
	}

	return results
}

func containsPath(paths [][]Position, path []Position) bool {
	if len(path) == 0 {
		return false
	}

	for _, existingPath := range paths {
		if len(existingPath) != len(path) {
			continue
		}

		matches := true
		for i := range path {
			if !path[i].Equals(existingPath[i]) {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}

func checkPaths(
	starting Position,
	dataPosition Position,
	data [][]int,
	completedPaths []Position,
) []Position {
	pathsInBounds := []Position{}
	for _, direction := range Directions[:4] {
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
			if !foundPath.IncludedIn(results) {
				results = append(results, foundPath)
			}
		}
	}

	return results
}

func (position Position) IncludedIn(positions []Position) bool {
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
