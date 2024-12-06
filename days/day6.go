package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	openSpot, obstacle    = byte(46), byte(35)
	up, down, left, right = byte(94), byte(118), byte(60), byte(62)
	directionMap          = map[byte]Direction{ // directions from day6
		left:  directions[0],
		right: directions[1],
		up:    directions[2],
		down:  directions[3],
	}
	directionSelectionMap = map[byte]byte{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}
)

type GuardRoute struct {
	Data []string
	// 	DirectionMap map[byte]Direction
	// 	Up byte
	// Down byte
	// 	Left byte
	// 	Right byte
	// 	OpenSpot byte
	// 	Obstacle byte
	GuardMap  map[string]bool
	Facing    byte
	Position  Direction
	ColLen    int
	RowHeight int
}

func Day6() {
	guardRoute := initializeDay6()

	for {
		direction := directionMap[guardRoute.Facing]
		newRow := guardRoute.Position.Row + direction.Row
		newCol := guardRoute.Position.Col + direction.Col

		isObstacle, end := checkIfObstruction(guardRoute, newRow, newCol)
		if end {
			break
		}

		if isObstacle {
			guardRoute.Facing = directionSelectionMap[guardRoute.Facing]
			continue
		}

		guardRoute.GuardMap[fmt.Sprintf("%d-%d", newRow, newCol)] = true
		guardRoute.Position.Row = newRow
		guardRoute.Position.Col = newCol
	}

	totalPositions := 0
	for range guardRoute.GuardMap {
		totalPositions++
	}

	log.Printf("Found %d Guard Postions", totalPositions)
}

func checkIfObstruction(guardRoute GuardRoute, newRow, newCol int) (bool, bool) {
	if !isInBounds(newRow, newCol, guardRoute.RowHeight, guardRoute.ColLen) {
		return false, true
	}

	if guardRoute.Data[newRow][newCol] == obstacle {
		return true, false
	}

	return false, false
}

func initializeDay6() GuardRoute {
	data := getDay6Data()
	var guardRoute GuardRoute
	guardRoute.Data = data
	guardRoute.GuardMap = make(map[string]bool)

	testArray := []byte{up, down, left, right}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if index := slices.Index(testArray, data[i][j]); index != -1 {
				guardRoute.Facing = testArray[index]
				guardRoute.Position.Row = i
				guardRoute.Position.Col = j
				guardRoute.GuardMap[fmt.Sprintf("%d-%d", i, j)] = true
			}
		}
	}

	guardRoute.RowHeight = len(data)
	guardRoute.ColLen = len(data[0])

	return guardRoute
}

func getDay6Data() []string {
	file, err := os.Open("inputs/input6.txt")
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
