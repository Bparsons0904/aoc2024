package days

import (
	"bufio"
	"log"
	"math"
	"os"

	"github.com/jinzhu/copier"
)

var (
	start = rune(83)
	end   = rune(69)
)

type PosMap map[Position]rune

type PosKey struct {
	Position Position
	Facing   Position
}

type Maze struct {
	Map      PosMap
	RowLen   int
	ColLen   int
	Start    Position
	End      Position
	Position Position
	Facing   Position
	Result   int
	Path     []PosKey
}

func Day16() {
	maze := getDay16Data()

	// maze.PrintState()
	lowestCount, sharedPathLen := traverseMap(maze)
	log.Printf("Final lowest count: %d with %d shared poisitions", lowestCount, sharedPathLen)
}

type QueueItem struct {
	Maze Maze
	Cost int
}

func traverseMap(maze Maze) (int, int) {
	queue := []QueueItem{{
		Maze: maze,
		Cost: maze.Result,
	}}
	lowestCount := math.MaxInt
	foundPaths := [][]PosKey{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Cost > lowestCount {
			continue
		}

		if current.Maze.Map[current.Maze.Position] == end {
			if current.Cost <= lowestCount {
				if current.Cost < lowestCount {
					lowestCount = current.Cost
					foundPaths = [][]PosKey{current.Maze.Path}
				} else {
					foundPaths = append(foundPaths, current.Maze.Path)
				}
			}
			continue
		}

		moves := []Position{current.Maze.Facing}
		if current.Maze.Facing.Row == 0 {
			moves = append(moves, Up, Down)
		} else {
			moves = append(moves, Left, Right)
		}

		possibleMoves := current.Maze.tryMove(moves)
		for _, newMaze := range possibleMoves {
			queue = append(queue, QueueItem{
				Maze: newMaze,
				Cost: newMaze.Result,
			})
		}
	}

	sharePositions := make(map[Position]bool)
	for _, paths := range foundPaths {
		for _, path := range paths {
			sharePositions[path.Position] = true
		}
	}

	return lowestCount, len(sharePositions)
}

func (maze *Maze) tryMove(moves []Position) []Maze {
	var possibleMoves []Maze
	for _, move := range moves {
		nextPosition := maze.Position.GetNextPosition(move)
		if nextPosition.IsWall(maze.Map) {
			continue
		}

		newMaze := Maze{}
		err := copier.Copy(&newMaze, maze)
		if err != nil {
			log.Panicf("Error trying to copy struct, %v", maze)
		}
		newPath := make([]PosKey, len(maze.Path))
		copy(newPath, maze.Path)
		newMaze.Path = newPath
		if move == maze.Facing {
			newMaze.Result += 1
		} else {
			newMaze.Result += 1001
		}
		newMaze.Position = nextPosition
		newMaze.Facing = move
		newMaze.Path = append(newMaze.Path, PosKey{nextPosition, move})
		possibleMoves = append(possibleMoves, newMaze)
	}
	return possibleMoves
}

func (position Position) GetNextPosition(direction Position) Position {
	newPosition := Position{
		Row: position.Row + direction.Row,
		Col: position.Col + direction.Col,
	}
	return newPosition
}

func (position Position) IsWall(posMap PosMap) bool {
	return posMap[position] == wall
}

func (position *Position) Move(direction Position) {
	position.Row = position.Row + direction.Row
	position.Col = position.Col + direction.Col
}

func getDay16Data() Maze {
	file, err := os.Open("inputs/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maze := Maze{
		Map:    make(PosMap),
		RowLen: 0,
		ColLen: 0,
		Facing: Right,
		Result: 0,
		Path:   make([]PosKey, 0),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()

		for j, val := range row {
			maze.Map[Position{maze.RowLen, j}] = val
			if val == start {
				maze.Position = Position{maze.RowLen, j}
				maze.Path = append(maze.Path, PosKey{maze.Position, Right})
			}
		}

		if maze.RowLen == 0 {
			maze.ColLen = len(row)
		}
		maze.RowLen++
	}

	return maze
}

func (maze Maze) PrintState() {
	for i, step := range maze.Path {
		log.Printf(
			"%d: Position(%d,%d) Facing:%v",
			i,
			step.Position.Row,
			step.Position.Col,
			step.Facing,
		)
	}

	pathMap := make(map[Position]rune)
	for i, step := range maze.Path {
		var dirChar rune
		switch step.Facing {
		case Up:
			dirChar = '^'
		case Down:
			dirChar = 'v'
		case Left:
			dirChar = '<'
		case Right:
			dirChar = '>'
		}

		pathMap[step.Position] = dirChar

		log.Printf("Adding to visualization: step %d at (%d,%d) char: %c",
			i, step.Position.Row, step.Position.Col, dirChar)
	}

	for row := 0; row < maze.RowLen; row++ {
		line := ""
		for col := 0; col < maze.ColLen; col++ {
			pos := Position{row, col}
			if pathChar, exists := pathMap[pos]; exists {
				line += string(pathChar)
			} else {
				line += string(maze.Map[pos])
			}
		}
		log.Println(line)
	}

	if len(maze.Path) > 1 {
		forwardMoves := 0
		turns := 0
		for i := 1; i < len(maze.Path); i++ {
			if maze.Path[i].Facing == maze.Path[i-1].Facing {
				forwardMoves++
			} else {
				turns++
			}
		}
		log.Printf("Stats: %d forwards (cost:%d) + %d turns (cost:%d) = total:%d",
			forwardMoves, forwardMoves,
			turns, turns*1000,
			forwardMoves+(turns*1000))
	}
}
