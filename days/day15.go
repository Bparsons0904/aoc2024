package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	space       = rune(46)
	wall        = rune(35)
	object      = rune(79)
	robot       = rune(64)
	leftObject  = rune('[')
	rightObject = rune(']')
	Up          = Position{-1, 0}
	Down        = Position{1, 0}
	Right       = Position{0, 1}
	Left        = Position{0, -1}
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
	moveWarehouseRobotExpanded(&warehouse)

	// gpsCordinates := calculateGPSCoordinates(warehouse)
	// gpsCordinatesExpanded := calculateGPSCoordinatesExpanded(warehouse)
	// log.Printf("Laternfish GPS Coordinates: %d ", gpsCordinates)
}

func moveWarehouseRobotExpanded(warehouse *Warehouse) {
	warehouse.RobotPosition = Position{warehouse.RobotStarting.Row, warehouse.RobotStarting.Col * 2}

	printWarehouseLayout(*warehouse, true)
	for i, movement := range warehouse.Movements {
		handleNextRobotMoveExpanded(warehouse, movement)
		printWarehouseLayout(*warehouse, true)
		if i > 6 {
			break
		}
	}
}

func handleNextRobotMoveExpanded(warehouse *Warehouse, movement Position) {
	newPosition := warehouse.RobotPosition.GetNextMove(movement)
	newSpace := warehouse.LayoutExpanded[newPosition]
	log.Println("Next move", newPosition, movement, newSpace)
	switch newSpace {
	case space:
		warehouse.LayoutExpanded[warehouse.RobotPosition] = space
		warehouse.LayoutExpanded[newPosition] = robot
		warehouse.RobotPosition = newPosition
	case leftObject:
		moveObjectExpanded(warehouse, newPosition, movement)
	case rightObject:
		moveObjectExpanded(warehouse, newPosition, movement)
	}
}

func moveObjectExpanded(warehouse *Warehouse, newPosition, movement Position) {
	switch movement {
	case Down:
		offset := 0
		if warehouse.LayoutExpanded[newPosition] == rightObject {
			offset = -1
		}
		leftTest := warehouse.LayoutExpanded[Position{newPosition.Row + 1, newPosition.Col + offset}]
		rightTest := warehouse.LayoutExpanded[Position{newPosition.Row + 1, newPosition.Col + offset + 1}]
		if leftTest == space && rightTest == space {
			log.Println("in move down")
			warehouse.LayoutExpanded[Position{newPosition.Row + 1, newPosition.Col + offset}] = leftObject
			warehouse.LayoutExpanded[Position{newPosition.Row + 1, newPosition.Col + offset + 1}] = rightObject
			warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col + offset}] = space
			warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col + offset + 1}] = space
			warehouse.LayoutExpanded[newPosition] = robot
			warehouse.LayoutExpanded[warehouse.RobotPosition] = space
			warehouse.RobotPosition = newPosition
			break
		}
		foundSpaceIndex := -1
		rowsChecked := 0
		hasEmptySpace := false
		for i := newPosition.Row + 1; i > warehouse.RowLen; i++ {
			spacesFound := 0
			wallFound := false
			for j := newPosition.Col + offset - 1 - rowsChecked; j <= newPosition.Col+offset+2+rowsChecked; j++ {
				testPosition := warehouse.LayoutExpanded[Position{i, j}]
				if testPosition == wall {
					wallFound = true
					break
				}
				if testPosition == leftObject || testPosition == rightObject {
					continue
				}
				spacesFound++
			}
			if wallFound {
				break
			}
			rowsChecked++
			foundSpaceIndex = i
			if spacesFound == (rowsChecked*2)+2 {
				hasEmptySpace = true
				break
			}

		}

		if foundSpaceIndex != -1 && hasEmptySpace {
			for i := foundSpaceIndex; i <= newPosition.Row; i++ {
				for j := newPosition.Col + offset - rowsChecked; j <= newPosition.Col+offset+1+rowsChecked; j++ {
					swapPosition := Position{i, j}
					warehouse.LayoutExpanded[swapPosition] = warehouse.LayoutExpanded[Position{swapPosition.Row - 1, swapPosition.Col}]
					warehouse.LayoutExpanded[Position{swapPosition.Row - 1, swapPosition.Col}] = space
				}
			}
			warehouse.RobotPosition = newPosition
		}
	case Up:
		offset := 0
		if warehouse.LayoutExpanded[newPosition] == rightObject {
			offset = -1
		}
		leftTest := warehouse.LayoutExpanded[Position{newPosition.Row - 1, newPosition.Col + offset}]
		rightTest := warehouse.LayoutExpanded[Position{newPosition.Row - 1, newPosition.Col + offset + 1}]
		if leftTest == space && rightTest == space {
			warehouse.LayoutExpanded[Position{newPosition.Row - 1, newPosition.Col + offset}] = leftObject
			warehouse.LayoutExpanded[Position{newPosition.Row - 1, newPosition.Col + offset + 1}] = rightObject
			warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col + offset}] = space
			warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col + offset + 1}] = space
			warehouse.LayoutExpanded[newPosition] = robot
			warehouse.LayoutExpanded[warehouse.RobotPosition] = space
			warehouse.RobotPosition = newPosition
			break
		}
		foundSpaceIndex := -1
		rowsChecked := 0
		hasEmptySpace := false
		for i := newPosition.Row - 1; i > 0; i-- {
			spacesFound := 0
			wallFound := false
			for j := newPosition.Col + offset - 1 - rowsChecked; j <= newPosition.Col+offset+2+rowsChecked; j++ {
				testPosition := warehouse.LayoutExpanded[Position{i, j}]
				if testPosition == wall {
					wallFound = true
					break
				}
				if testPosition == leftObject || testPosition == rightObject {
					continue
				}
				spacesFound++
			}
			if wallFound {
				break
			}
			rowsChecked++
			foundSpaceIndex = i
			if spacesFound == (rowsChecked*2)+2 {
				hasEmptySpace = true
				break
			}

		}

		if foundSpaceIndex != -1 && hasEmptySpace {
			for i := foundSpaceIndex; i <= newPosition.Row; i++ {
				for j := newPosition.Col + offset - rowsChecked; j <= newPosition.Col+offset+1+rowsChecked; j++ {
					swapPosition := Position{i, j}
					warehouse.LayoutExpanded[swapPosition] = warehouse.LayoutExpanded[Position{swapPosition.Row + 1, swapPosition.Col}]
					warehouse.LayoutExpanded[Position{swapPosition.Row + 1, swapPosition.Col}] = space
				}
			}
			warehouse.RobotPosition = newPosition
		}

	case Right:
		for i := newPosition.Col + 1; i < warehouse.ColLen; i++ {
			testPosition := warehouse.LayoutExpanded[Position{newPosition.Row, i}]
			if testPosition == wall {
				break
			}
			if testPosition == space {
				log.Println("Found a test positions")
				for j := i; j > 0; j-- {
					warehouse.LayoutExpanded[Position{newPosition.Row, j}] = warehouse.LayoutExpanded[Position{newPosition.Row, j - 1}]
				}
				warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col - 1}] = space
				warehouse.RobotPosition = Position{newPosition.Row, newPosition.Col}
				break
			}
		}

	case Left:
		for i := newPosition.Col - 1; i > 0; i-- {
			testPosition := warehouse.LayoutExpanded[Position{newPosition.Row, i}]
			if testPosition == wall {
				break
			}
			if testPosition == space {
				for j := i; j <= newPosition.Col; j++ {
					warehouse.LayoutExpanded[Position{newPosition.Row, j}] = warehouse.LayoutExpanded[Position{newPosition.Row, j + 1}]
				}
				warehouse.LayoutExpanded[Position{newPosition.Row, newPosition.Col + 1}] = space
				warehouse.RobotPosition = Position{newPosition.Row, newPosition.Col}
				break
			}
		}
	}
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

func printWarehouseLayout(warehouse Warehouse, isExpanded bool) {
	layout := warehouse.LayoutExpanded
	colLen, rowLen := warehouse.ColLen*2, warehouse.RowLen
	if !isExpanded {
		// layout = warehouse.LayoutExpanded
		// rowLen = warehouse.RowLen
		return
	}

	result := ""
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			result += string(layout[Position{i, j}])
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
	case object:
		pushObject(warehouse, movement, newPosition)
	}
}

func pushObject(warehouse *Warehouse, movement Position, objectPosition Position) {
	positionPastObject := objectPosition.GetNextMove(movement)
	spacePastObject := warehouse.Layout[positionPastObject]
	switch spacePastObject {
	case space:
		warehouse.Layout[positionPastObject] = object
		newRobotPosition := warehouse.RobotPosition.GetNextMove(movement)
		warehouse.Layout[newRobotPosition] = robot
		warehouse.Layout[warehouse.RobotPosition] = space
		warehouse.RobotPosition = newRobotPosition
	case object:
		pushObject(warehouse, movement, positionPastObject)
	}
}

func (position Position) GetNextMove(direction Position) Position {
	return Position{position.Row + direction.Row, position.Col + direction.Col}
}

func getDay15Data() Warehouse {
	file, err := os.Open("inputs/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	warehouse := Warehouse{
		Layout:         make(map[Position]rune),
		LayoutExpanded: make(map[Position]rune),
		Movements:      []Position{},
		RowLen:         0,
		ColLen:         0,
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
				switch location {
				case wall:
					warehouse.Layout[Position{warehouse.RowLen, col}] = location
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col * 2}] = location
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col*2 + 1}] = location
				case space:
					warehouse.Layout[Position{warehouse.RowLen, col}] = location
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col * 2}] = location
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col*2 + 1}] = location
				case object:
					warehouse.Layout[Position{warehouse.RowLen, col}] = object
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col * 2}] = leftObject
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col*2 + 1}] = rightObject
				case robot:
					warehouse.Layout[Position{warehouse.RowLen, col}] = robot
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col * 2}] = robot
					warehouse.LayoutExpanded[Position{warehouse.RowLen, col*2 + 1}] = space
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
