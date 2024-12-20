package days

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	log.Println(memory.Map)
	memorySpaces := calculateMemoryEscape(memory)
	log.Println(memorySpaces)
}

type MemoryPath struct {
	Path []Position
}

type PathState struct {
	Pos      Position
	Distance int
}

func calculateMemoryEscape(baseMemory Memory) int {
	queue := []PathState{{Pos: Position{0, 0}, Distance: 0}}
	visited := make(map[Position]int) // Position -> shortest distance
	visited[Position{0, 0}] = 0
	shortestPath := math.MaxInt
	endPosition := Position{baseMemory.RowLen, baseMemory.ColLen}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we found end position, update shortest path if needed
		if current.Pos == endPosition {
			if current.Distance < shortestPath {
				shortestPath = current.Distance
			}
			continue
		}

		// Check all possible moves
		for _, direction := range Directions[:4] {
			newPos := Position{
				Row: current.Pos.Row + direction.Row,
				Col: current.Pos.Col + direction.Col,
			}

			// Skip if out of bounds
			if !isInBounds(newPos.Row, newPos.Col, baseMemory.RowLen+1, baseMemory.ColLen+1) {
				continue
			}

			// Skip if position is corrupted
			if byteDrop, ok := baseMemory.Map[newPos]; ok && byteDrop < 1024 {
				continue
			}

			// Skip if we've found a shorter path to this position already
			if prevDist, exists := visited[newPos]; exists && current.Distance+1 >= prevDist {
				continue
			}

			// Add new position to queue and update visited map
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

func getPossiblePaths(
	baseMemory Memory,
	memory MemoryPath,
	memoryLen int,
) []Position {
	possiblePositions := []Position{}
	for _, direction := range Directions[:4] {
		newPath := memory.GetNextPosition(direction)
		inBounds := isInBounds(newPath.Row, newPath.Col, memoryLen+1, memoryLen+1)
		if !inBounds {
			continue
		}

		byteDrop, ok := baseMemory.Map[newPath]
		if ok && byteDrop < 1024 {
			continue
		}

		if slices.ContainsFunc(memory.Path, func(prevPath Position) bool {
			return prevPath.Equals(newPath)
		}) {
			continue
		}

		possiblePositions = append(possiblePositions, newPath)

	}

	return possiblePositions
}

func printMemoryVisualization(memory Memory, nanosecond int) {
	result := ""
	for i := 0; i <= memory.RowLen; i++ {
		for j := 0; j <= memory.ColLen; j++ {
			memoryPosition, ok := memory.Map[Position{i, j}]
			if !ok {
				result += string(space)
				continue
			}
			if memoryPosition <= nanosecond {
				result += "#"
			} else {
				result += string(space)
			}
		}
		result += "\n"
	}

	fmt.Println(result)
}

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
