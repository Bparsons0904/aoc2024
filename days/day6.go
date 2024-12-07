package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"
)

var (
	obstacle              = byte(35)
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
	Data        []string
	ParadoxLoop []Direction
	GuardMap    map[string]bool
	Facing      byte
	Position    Direction
	ColLen      int
	RowHeight   int
}

func Day6() {
	data := getDay6Data()
	var guardRoute GuardRoute
	guardRoute.Data = data
	guardRoute.GuardMap = make(map[string]bool)
	guardRoute.ParadoxLoop = []Direction{}
	initializeDay6(&guardRoute, data)

	totalPositions := getGuardPositions(&guardRoute)
	totalParadoxes := getGuardParadoxes(guardRoute)

	log.Printf(
		"Found %d Guard Postions and Found %d Paradox Loops",
		totalPositions,
		totalParadoxes,
	)
}

func getGuardParadoxes(guardRoute GuardRoute) int {
	var startRow, startCol int
	for i := 0; i < guardRoute.RowHeight; i++ {
		for j := 0; j < guardRoute.ColLen; j++ {
			if guardRoute.Data[i][j] == up || guardRoute.Data[i][j] == down ||
				guardRoute.Data[i][j] == left || guardRoute.Data[i][j] == right {
				startRow, startCol = i, j
				break
			}
		}
	}

	results := make(chan int)
	var wg sync.WaitGroup

	checkPosition := func(i, j int) {
		defer wg.Done()

		if guardRoute.Data[i][j] == obstacle ||
			(i == startRow && j == startCol) ||
			!guardRoute.GuardMap[fmt.Sprintf("%d-%d", i, j)] {
			return
		}

		testGuard := GuardRoute{
			Data:      guardRoute.Data,
			GuardMap:  make(map[string]bool),
			RowHeight: guardRoute.RowHeight,
			ColLen:    guardRoute.ColLen,
		}
		initializeDay6(&testGuard, guardRoute.Data)

		visitedStates := make(map[string]bool)
		firstLoop := true

		for {
			direction := directionMap[testGuard.Facing]
			newRow := testGuard.Position.Row + direction.Row
			newCol := testGuard.Position.Col + direction.Col

			isObstacle, end := checkIfObstruction(testGuard, newRow, newCol)
			if end {
				break
			}

			stateKey := fmt.Sprintf("%d-%d-%d",
				testGuard.Position.Row,
				testGuard.Position.Col,
				testGuard.Facing)

			if !firstLoop && visitedStates[stateKey] {
				results <- 1
				return
			}

			visitedStates[stateKey] = true

			if isObstacle || (newRow == i && newCol == j) {
				testGuard.Facing = directionSelectionMap[testGuard.Facing]
				firstLoop = false
				continue
			}

			testGuard.Position.Row = newRow
			testGuard.Position.Col = newCol
		}
		results <- 0
	}

	for i := 0; i < guardRoute.RowHeight; i++ {
		for j := 0; j < guardRoute.ColLen; j++ {
			wg.Add(1)
			go checkPosition(i, j)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	paradoxLoops := 0
	for result := range results {
		paradoxLoops += result
	}

	return paradoxLoops
}

func getGuardPositions(guardRoute *GuardRoute) int {
	for {
		direction := directionMap[guardRoute.Facing]
		newRow := guardRoute.Position.Row + direction.Row
		newCol := guardRoute.Position.Col + direction.Col

		isObstacle, end := checkIfObstruction(*guardRoute, newRow, newCol)
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

		guardRoute.ParadoxLoop = append(guardRoute.ParadoxLoop, guardRoute.Position)
	}

	totalPositions := 0
	for range guardRoute.GuardMap {
		totalPositions++
	}

	return totalPositions
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

func initializeDay6(guardRoute *GuardRoute, data []string) {
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
