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
	A, B, Prize, PrizeExtended Cord
	Coins                      int
	CoinsExtended              int64
}

var clawPrizePattern = regexp.MustCompile(`([AB]|Prize): X=?([+-]?\d+), Y=?([+-]?\d+)`)

func Day13() {
	data := getDay13Data()
	log.Println("Getitng data")
	fewestTokens, fewestTokensExtend := calculateTokenCounts(data)
	log.Printf(
		"The total coins used for the most games played is %d and with values extended %d",
		fewestTokens,
		fewestTokensExtend,
	)
}

func calculateTokenCounts(clawMachines []ClawMachine) (int, int64) {
	result := 0
	resultExtended := int64(0)

	for _, clawMachine := range clawMachines {
		calculateOptions(&clawMachine)
		result += clawMachine.Coins
		calculateOptionsExtended(&clawMachine)
		resultExtended += clawMachine.CoinsExtended
	}
	return result, resultExtended
}

func calculateOptionsExtended(clawMachine *ClawMachine) {
	const UNIT_CONVERSION int64 = 10000000000000

	aX := int64(clawMachine.A.X)
	aY := int64(clawMachine.A.Y)
	bX := int64(clawMachine.B.X)
	bY := int64(clawMachine.B.Y)
	xExpanded := int64(clawMachine.Prize.X) + UNIT_CONVERSION
	yExpanded := int64(clawMachine.Prize.Y) + UNIT_CONVERSION

	determinant := aX*bY - bX*aY
	determiniant1 := xExpanded*bY - yExpanded*bX
	determinant2 := yExpanded*aX - xExpanded*aY

	if determiniant1%determinant != 0 || determinant2%determinant != 0 {
		return
	}

	solution1 := determiniant1 / determinant
	solution2 := determinant2 / determinant

	if solution1 > 0 && solution2 > 0 {
		tokens := 3*solution1 + solution2
		clawMachine.CoinsExtended = tokens
	}
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

		matches := clawPrizePattern.FindStringSubmatch(line)
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
