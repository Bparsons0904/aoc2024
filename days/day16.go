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
	Visited  map[PosKey]int
	Result   int
	Path     []PosKey
}

func Day16() {
	maze := getDay16Data()

	maze.PrintState()
	traverseMap(maze)
}

func traverseMap(maze Maze) {
	toTraverse := []Maze{maze}
	lowestCount := math.MaxInt

	for len(toTraverse) > 0 {
		idx := len(toTraverse) - 1
		traversing := toTraverse[idx]
		toTraverse = toTraverse[:idx]

		if traversing.Map[traversing.Position] == end {
			if traversing.Result < lowestCount {
				lowestCount = traversing.Result
				log.Printf("Found new lowest count: %d", lowestCount)
				traversing.PrintState()
			}
			continue
		}

		if traversing.Result >= lowestCount {
			continue
		}

		var moves []Position
		switch traversing.Facing {
		case Up:
			moves = []Position{Up, Left, Right}
		case Down:
			moves = []Position{Down, Left, Right}
		case Left:
			moves = []Position{Left, Up, Down}
		case Right:
			moves = []Position{Right, Up, Down}
		}

		possibleMoves := traversing.tryMove(moves)
		toTraverse = append(toTraverse, possibleMoves...)
	}

	if lowestCount == math.MaxInt {
		log.Println("No path found!")
	} else {
		log.Printf("Final lowest count: %d", lowestCount)
	}
}

func (maze Maze) PrintState() {
	log.Println("Path sequence:")
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
	log.Println("-------------------")
}

func (maze *Maze) tryMove(moves []Position) []Maze {
	var possibleMoves []Maze
	for _, move := range moves {
		nextPosition := maze.Position.GetNextPosition(move)

		rowDiff := abs(nextPosition.Row - maze.Position.Row)
		colDiff := abs(nextPosition.Col - maze.Position.Col)
		if rowDiff+colDiff != 1 {
			log.Printf("Skip discontinuous move from (%d,%d) to (%d,%d)",
				maze.Position.Row, maze.Position.Col,
				nextPosition.Row, nextPosition.Col)
			continue
		}

		if nextPosition.IsWall(maze.Map) {
			log.Printf("Skip wall at (%d,%d)", nextPosition.Row, nextPosition.Col)
			continue
		}

		visitKey := PosKey{
			Position: nextPosition,
			Facing:   move,
		}

		if score, exists := maze.Visited[visitKey]; exists && score <= maze.Result {
			log.Printf("Skip visited state at (%d,%d) facing %v with better score %d vs current %d",
				nextPosition.Row, nextPosition.Col, move, score, maze.Result)
			continue
		}

		newMaze := Maze{}
		err := copier.Copy(&newMaze, maze)
		if err != nil {
			log.Panicf("Error trying to copy struct, %v", maze)
		}

		oldScore := newMaze.Result
		if move == maze.Facing {
			newMaze.Result += 1
			log.Printf("Forward move: +1")
		} else {
			newMaze.Result += 1000
			log.Printf("Turn: +1000")
		}
		log.Printf("Score changed: %d -> %d", oldScore, newMaze.Result)

		newMaze.Visited[visitKey] = newMaze.Result
		newMaze.Position = nextPosition
		newMaze.Facing = move
		newMaze.Path = append(newMaze.Path, PosKey{nextPosition, move})

		possibleMoves = append(possibleMoves, newMaze)
	}
	return possibleMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (position Position) GetNextPosition(direction Position) Position {
	return Position{
		Row: position.Row + direction.Row,
		Col: position.Col + direction.Col,
	}
}

func (position Position) IsWall(posMap PosMap) bool {
	return posMap[position] == wall
}

func (position *Position) Move(direction Position) {
	position.Row = position.Row + direction.Row
	position.Col = position.Col + direction.Col
}

// func (position Position) IsWall(direction Position, posMap PosMap, rowLen, colLen int) bool {
// 	newPos := Position{
// 		Row: position.Row + direction.Row,
// 		Col: position.Col + direction.Col,
// 	}
//
// 	// Check boundaries first
// 	if newPos.Row < 0 || newPos.Row >= rowLen ||
// 		newPos.Col < 0 || newPos.Col >= colLen {
// 		return true // Out of bounds is treated as a wall
// 	}
//
// 	return posMap[newPos] == wall
// }

func getDay16Data() Maze {
	file, err := os.Open("inputs/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maze := Maze{
		Map:     make(PosMap),
		Visited: make(map[PosKey]int),
		RowLen:  0,
		ColLen:  0,
		Facing:  Right,
		Result:  0,
		Path:    make([]PosKey, 0),
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

// forwardPosition := traversing.Position.GetNextMove(Right)
// if !forwardPosition.IsWall(Right, traversing.Map) &&
// 	!traversing.Visited[forwardPosition] {
// 	newMaze := Maze{}
// 	copier.Copy(&newMaze, &traversing)
// 	newMaze.Visited[forwardPosition] = true
// 	newMaze.Result++
// 	newMaze.Position.Move(Right)
// 	toTraverse = append(toTraverse, newMaze)
// }
//
// leftPosition := traversing.Position.GetNextMove(Up)
// if !leftPosition.IsWall(Up, traversing.Map) &&
// 	!traversing.Visited[leftPosition] {
// 	newMaze := Maze{}
// 	copier.Copy(&newMaze, &traversing)
// 	newMaze.Visited[leftPosition] = true
// 	newMaze.Result += 90
// 	newMaze.Position.Move(Up)
// 	newMaze.Facing = Up
// 	toTraverse = append(toTraverse, newMaze)
// }
//
// rightPosition := traversing.Position.GetNextMove(Down)
// if !rightPosition.IsWall(Down, traversing.Map) &&
// 	!traversing.Visited[rightPosition] {
// 	newMaze := Maze{}
// 	copier.Copy(&newMaze, &traversing)
// 	newMaze.Visited[rightPosition] = true
// 	newMaze.Result += 90
// 	newMaze.Position.Move(Down)
// 	newMaze.Facing = Down
// 	toTraverse = append(toTraverse, newMaze)
// 	break
// }
