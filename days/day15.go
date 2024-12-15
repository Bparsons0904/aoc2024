package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	space  = rune(46)
	wall   = rune(35)
	object = rune(79)
	robot  = rune(64)
)

type Warehouse struct {
	Layout         map[Position]rune
	LayoutExpanded map[Position]rune
	Movements      []Position
	RowLen         int
	ColLen         int
	RobotStarting  Position
	RobotPosition  Position
}

func Day15() {
	warehouse := getDay15Data()
	moveWarehouseRobot(&warehouse)

	gpsCordinates := calculateGPSCoordinates(warehouse)
	log.Printf("Laternfish GPS Coordinates: %d ", gpsCordinates)
}

func calculateGPSCoordinates(warehouse Warehouse) int {
	results := 0
	for key, value := range warehouse.Layout {
		if value == object {
			result := (key.Row * 100) + key.Col
			results += result
		}
	}

	return results
}

func moveWarehouseRobot(warehouse *Warehouse) {
	for _, movement := range warehouse.Movements {
		handleNextRobotMove(warehouse, movement)
	}
}

func printWarehouseLayout(warehouse Warehouse) {
	return
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
		warehouse.Layout[warehouse.RobotPosition] = space
		warehouse.Layout[newPosition] = robot
		warehouse.RobotPosition = newPosition
		printWarehouseLayout(*warehouse)
	case object:
		pushObject(warehouse, movement, newPosition)
	}
}

func pushObject(warehouse *Warehouse, movement Position, objectPosition Position) {
	positionPastObject := objectPosition.GetNextMove(movement)
	spacePastObject := warehouse.Layout[positionPastObject]
	switch spacePastObject {
	case space:
		printWarehouseLayout(*warehouse)
		warehouse.Layout[positionPastObject] = object
		newRobotPosition := warehouse.RobotPosition.GetNextMove(movement)
		warehouse.Layout[newRobotPosition] = robot
		warehouse.Layout[warehouse.RobotPosition] = space
		warehouse.RobotPosition = newRobotPosition
		printWarehouseLayout(*warehouse)
	case object:
		pushObject(warehouse, movement, positionPastObject)
	}
}

func (position Position) GetNextMove(direction Position) Position {
	return Position{position.Row + direction.Row, position.Col + direction.Col}
}

func getDay15Data() Warehouse {
	file, err := os.Open("inputs/input15.txt")
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
			for col, location := range row {
				warehouse.Layout[Position{warehouse.RowLen, col}] = location
				if location == robot {
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
