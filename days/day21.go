package days

import (
	"bufio"
	"log"
	"os"
)

var (
	keypadMap = map[byte]Position{
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
		'0': {3, 1},
		'A': {3, 2},
	}

	upArrow    = byte('^')
	downArrow  = byte('v')
	leftArrow  = byte('<')
	rightArrow = byte('>')
)

func Day21() {
	data := getDay21Data()
	log.Println(data)
	pressKeyPad(Position{1, 2}, Position{3, 2}, data[0][0])
}

func pressKeyPad(dPadCurrent, keypadCurrent Position, target byte) {
	log.Println(dPadCurrent, keypadCurrent, keypadMap[target])
	distanceToMove := keypadCurrent.DistanceToMove(keypadMap[target])

	presses := []byte{}
	log.Println("initial", distanceToMove)
	for !distanceToMove.Equals(Position{0, 0}) {
		switch {
		case distanceToMove.Row < 0:
			log.Println("Move Down")
			presses = append(presses, downArrow)
			distanceToMove = distanceToMove.MoveTo(Down)
		case distanceToMove.Row > 0:
			log.Println("Move Up")
			presses = append(presses, upArrow)
			distanceToMove = distanceToMove.Shift(Up)

		case distanceToMove.Col < 0:
			log.Println("Move Left")
			presses = append(presses, leftArrow)
			distanceToMove = distanceToMove.Shift(Left)
		case distanceToMove.Col > 0:
			log.Println("Move Right")
			presses = append(presses, rightArrow)
			distanceToMove = distanceToMove.Shift(Right)
		}
		log.Println("after move", distanceToMove)
	}
	log.Println(presses)
}

func (position Position) Shift(shift Position) Position {
	newRow, newCol := position.Row-shift.Row, position.Col-shift.Col
	return Position{
		newRow, newCol,
	}
}

func (position Position) DistanceToMove(moveTo Position) Position {
	return Position{moveTo.Row - position.Row, moveTo.Col - position.Col}
}

func getDay21Data() []string {
	file, err := os.Open("inputs/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		// num, err := strconv.Atoi(row[:len(row)-1])
		// if err != nil {
		// 	log.Fatal(err)
		// }
		results = append(results, row[:len(row)-1])
	}

	return results
}
