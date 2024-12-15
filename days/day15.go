package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// type Movement Position
//
// type Movement struct{
//  Up Position
//
// Down	Position
//  Left Position
//  Right Position
//
// }

type Warehouse struct {
	Layout        map[Position]rune
	Movements     []Position
	RowLen        int
	ColLen        int
	RobotStarting Position
	RobotPosition Position
}

func Day15() {
	warehouse := getDay15Data()
	log.Println(warehouse)
	moveWarehouseRobot(&warehouse)
}

var (
	space  = rune(46)
	wall   = rune(35)
	object = rune(79)
	robot  = rune(64)
)

func moveWarehouseRobot(warehouse *Warehouse) {
	for i, movement := range warehouse.Movements {
		log.Println("Starting movements", i, movement)
		handleNextRobotMove(warehouse, movement)
		// if i > 10 {
		// 	break
		// }
	}
}

func printWarehouseLayout(warehouse Warehouse) {
	result := ""
	for i := 0; i < warehouse.RowLen; i++ {
		for j := 0; j < warehouse.ColLen; j++ {
			result += string(warehouse.Layout[Position{i, j}])
		}
		result += "\n"
	}

	fmt.Println(result)
}

func handleNextRobotMove(warehouse *Warehouse, movement Position) {
	newPosition := warehouse.RobotPosition.GetNextMove(movement)
	newSpace := warehouse.Layout[newPosition]
	switch newSpace {
	case space:
		log.Println("We have an open space", newPosition)
		warehouse.Layout[warehouse.RobotPosition] = space
		warehouse.Layout[newPosition] = robot
		warehouse.RobotPosition = newPosition
		printWarehouseLayout(*warehouse)
		// handleNextRobotMove(warehouse, movement)
	// case wall:
	// 	log.Println("We have an wall space")
	case object:
		log.Println("We have an object space")
		pushObject(warehouse, movement, newPosition)

	}
	// printWarehouseLayout(*warehouse)
}

func pushObject(warehouse *Warehouse, movement Position, objectPosition Position) {
	positionPastObject := objectPosition.GetNextMove(movement)
	spacePastObject := warehouse.Layout[positionPastObject]
	switch spacePastObject {
	case space:
		log.Println("We have an open space when trying to push", objectPosition)
		printWarehouseLayout(*warehouse)
		warehouse.Layout[positionPastObject] = object
		newRobotPosition := warehouse.RobotPosition.GetNextMove(movement)
		warehouse.Layout[newRobotPosition] = robot
		warehouse.Layout[warehouse.RobotPosition] = space
		warehouse.RobotPosition = newRobotPosition
		// Change this to a swap
		printWarehouseLayout(*warehouse)

		// pushObject(warehouse, movement, newPosition)
		// warehouse.RobotPosition = newPosition
	case wall:
		log.Println("We have an wall space trying to push")
	case object:
		pushObject(warehouse, movement, positionPastObject)

		log.Println("We have an object space trying to push")
	}
}

func (position Position) GetNextMove(direction Position) Position {
	return Position{position.Row + direction.Row, position.Col + direction.Col}
}

func (position *Position) Move(direction Position) {
	position.Row = position.Row + direction.Row
	position.Col = position.Col + direction.Col
}

func getDay15Data() Warehouse {
	file, err := os.Open("inputs/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	warehouse := Warehouse{
		Layout:    make(map[Position]rune),
		Movements: []Position{},
		RowLen:    0,
		ColLen:    0,
	}

	emptySpaceFound := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {
			emptySpaceFound = true
			continue
		}

		switch emptySpaceFound {
		case false:
			for col, space := range row {
				warehouse.Layout[Position{warehouse.RowLen, col}] = space
				if space == 64 { // Robot is a @
					warehouse.RobotStarting = Position{warehouse.RowLen, col}
					warehouse.RobotPosition = Position{warehouse.RowLen, col}
				}
			}
			warehouse.RowLen++
			warehouse.ColLen = len(row)
		case true:
			for _, direction := range row {
				switch direction {
				case 60: // Left
					warehouse.Movements = append(warehouse.Movements, directions[0])
				case 62: // Right
					warehouse.Movements = append(warehouse.Movements, directions[1])
				case 94: // Up
					warehouse.Movements = append(warehouse.Movements, directions[2])
				case 118: // Down
					warehouse.Movements = append(warehouse.Movements, directions[3])
				}
			}
		}

	}

	return warehouse
}
