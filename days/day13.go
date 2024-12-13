package days

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Cord struct {
	X, Y int
}
type ClawMachine struct {
	A, B, Prize Cord
	Coins       int
}

var coordPattern = regexp.MustCompile(`([AB]|Prize): X=?([+-]?\d+), Y=?([+-]?\d+)`)

func Day13() {
	data := getDay13Data()
	fewestTokens := calculateTokenCounts(data)
	log.Printf("The total coins used for the most games played is %d", fewestTokens)
}

func calculateTokenCounts(clawMachines []ClawMachine) int {
	result := 0
	for _, clawMachine := range clawMachines {
		// if i < 2 {
		calculateOptions(&clawMachine)
		result += clawMachine.Coins
		// }
	}
	return result
}

func calculateOptions(clawMachine *ClawMachine) {
	for a := 1; a <= 100; a++ {
		currentX := clawMachine.Prize.X - a*clawMachine.A.X
		currentY := clawMachine.Prize.Y - a*clawMachine.A.Y

		if currentX < 0 || currentY < 0 {
			break
		}

		if currentX%clawMachine.B.X == 0 && currentY%clawMachine.B.Y == 0 {
			bMovesForX := currentX / clawMachine.B.X
			bMovesForY := currentY / clawMachine.B.Y

			if bMovesForX == bMovesForY {
				tempACount := a
				tempBCount := bMovesForX
				tempCoinsUsed := (tempACount * 3) + tempBCount

				if (clawMachine.Coins == 0 || tempCoinsUsed < clawMachine.Coins) &&
					tempBCount <= 100 {
					clawMachine.Coins = tempCoinsUsed
				}
			}
		}
	}

	for b := 1; b <= 100; b++ {
		currentX := clawMachine.Prize.X - b*clawMachine.B.X
		currentY := clawMachine.Prize.Y - b*clawMachine.B.Y

		if currentX < 0 || currentY < 0 {
			break
		}

		if currentX%clawMachine.A.X == 0 && currentY%clawMachine.A.Y == 0 {
			aMovesForX := currentX / clawMachine.A.X
			aMovesForY := currentY / clawMachine.A.Y

			if aMovesForX == aMovesForY {
				tempACount := aMovesForX
				tempBCount := b
				tempCoinsUsed := (tempACount * 3) + tempBCount

				if (clawMachine.Coins == 0 || tempCoinsUsed < clawMachine.Coins) &&
					tempACount <= 100 {
					clawMachine.Coins = tempCoinsUsed
				}
			}
		}
	}
}

func getDay13Data() []ClawMachine {
	file, err := os.Open("inputs/input13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	clawMachines := []ClawMachine{}
	scanner := bufio.NewScanner(file)

	var currentMachine ClawMachine

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		matches := coordPattern.FindStringSubmatch(line)
		if len(matches) == 4 {
			x, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Panicln("Error parsing x int", x, err)
			}
			y, err := strconv.Atoi(matches[3])
			if err != nil {
				log.Panicln("Error parsing y int", y, err)
			}

			switch matches[1] {
			case "A":
				currentMachine.A = Cord{X: x, Y: y}
			case "B":
				currentMachine.B = Cord{X: x, Y: y}
			case "Prize":
				currentMachine.Prize = Cord{X: x, Y: y}
				clawMachines = append(clawMachines, currentMachine)
				currentMachine = ClawMachine{}
			}
		}
	}

	return clawMachines
}
