package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	Map    map[Position]int
	Order  []Position
	RowLen int
	ColLen int
}

func Day18() {
	memory := getDay18Data()
	memorySpaces := calculateMemoryEscape(memory, 1024)
	blockingPosition := findBlockingPath(memory)
	log.Printf(
		"Shortest path out out of the memory %d with a last blocking position of %s",
		memorySpaces,
		blockingPosition,
	)
}

type MemoryPath struct {
	Path []Position
}

type PathState struct {
	Pos      Position
	Distance int
}

func findBlockingPath(baseMemory Memory) string {
	low := 0
	high := 12
	lastFound := -1

	for {
		pathLength := calculateMemoryEscape(baseMemory, high)
		if pathLength == math.MaxInt {
			break
		}
		lastFound = high
		high *= 2
	}

	for low <= high {
		mid := low + (high-low)/2
		pathLength := calculateMemoryEscape(baseMemory, mid)

		if pathLength == math.MaxInt {
			high = mid - 1
		} else {
			lastFound = mid
			low = mid + 1
		}
	}

	blockingPosition := baseMemory.Order[lastFound]
	return fmt.Sprintf(
		"Minimum blocking value: {%d,%d}",
		blockingPosition.Col,
		blockingPosition.Row,
	)
}

func calculateMemoryEscape(baseMemory Memory, maxByteDrop int) int {
	queue := []PathState{{Pos: Position{0, 0}, Distance: 0}}
	visited := make(map[Position]int)
	visited[Position{0, 0}] = 0
	shortestPath := math.MaxInt
	endPosition := Position{baseMemory.RowLen, baseMemory.ColLen}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Pos == endPosition {
			if current.Distance < shortestPath {
				shortestPath = current.Distance
			}
			continue
		}

		for _, direction := range Directions[:4] {
			newPos := Position{
				Row: current.Pos.Row + direction.Row,
				Col: current.Pos.Col + direction.Col,
			}

			if !isInBounds(newPos.Row, newPos.Col, baseMemory.RowLen+1, baseMemory.ColLen+1) {
				continue
			}

			if byteDrop, ok := baseMemory.Map[newPos]; ok && byteDrop < maxByteDrop {
				continue
			}

			if prevDist, exists := visited[newPos]; exists && current.Distance+1 >= prevDist {
				continue
			}

			visited[newPos] = current.Distance + 1
			queue = append(queue, PathState{
				Pos:      newPos,
				Distance: current.Distance + 1,
			})
		}
	}

	return shortestPath
}

func (memoryPath MemoryPath) GetNextPosition(direction Position) Position {
	lastPosition := memoryPath.Path[len(memoryPath.Path)-1]
	return Position{lastPosition.Row + direction.Row, lastPosition.Col + direction.Col}
}

// func getPossiblePaths(
// 	baseMemory Memory,
// 	memory MemoryPath,
// 	memoryLen int,
// ) []Position {
// 	possiblePositions := []Position{}
// 	for _, direction := range Directions[:4] {
// 		newPath := memory.GetNextPosition(direction)
// 		inBounds := isInBounds(newPath.Row, newPath.Col, memoryLen+1, memoryLen+1)
// 		if !inBounds {
// 			continue
// 		}
//
// 		byteDrop, ok := baseMemory.Map[newPath]
// 		if ok && byteDrop < 12 {
// 			continue
// 		}
//
// 		if slices.ContainsFunc(memory.Path, func(prevPath Position) bool {
// 			return prevPath.Equals(newPath)
// 		}) {
// 			continue
// 		}
//
// 		possiblePositions = append(possiblePositions, newPath)
//
// 	}
//
// 	return possiblePositions
// }

// func printMemoryVisualization(memory Memory, nanosecond int) {
// 	result := ""
// 	for i := 0; i <= memory.RowLen; i++ {
// 		for j := 0; j <= memory.ColLen; j++ {
// 			memoryPosition, ok := memory.Map[Position{i, j}]
// 			if !ok {
// 				result += string(space)
// 				continue
// 			}
// 			if memoryPosition <= nanosecond {
// 				result += "#"
// 			} else {
// 				result += string(space)
// 			}
// 		}
// 		result += "\n"
// 	}
//
// 	fmt.Println(result)
// }

func getDay18Data() Memory {
	file, err := os.Open("inputs/input18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	memory := Memory{
		Map:   make(map[Position]int),
		Order: []Position{},
	}

	scanner := bufio.NewScanner(file)
	order := 0
	for scanner.Scan() {
		xy := strings.Split(scanner.Text(), ",")
		xPos, err := strconv.Atoi(string(xy[0]))
		if err != nil {
			log.Fatal(err)
		}
		yPos, err := strconv.Atoi(string(xy[1]))
		if err != nil {
			log.Fatal(err)
		}

		memory.Map[Position{yPos, xPos}] = order
		memory.Order = append(memory.Order, Position{yPos, xPos})
		if yPos > memory.RowLen {
			memory.RowLen = yPos
		}
		if xPos > memory.ColLen {
			memory.ColLen = xPos
		}
		order++

	}

	return memory
}
