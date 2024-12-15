package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var robotGuardPattern = regexp.MustCompile(`p=(\-?\d+),(\-?\d+) v=(\-?\d+),(\-?\d+)`)

type RobotGuard struct {
	Position Position
	Movement Position
}

type RobotGuardMap struct {
	MapXLength int
	MapYLength int
	Filename   string
	Time       int
}

func Day14() {
	// test := RobotGuardMap{
	// 	MapXLength: 11,
	// 	MapYLength: 7,
	// 	Filename:   "inputs/test.txt",
	// 	Time:       100,
	// }
	day14 := RobotGuardMap{
		MapXLength: 101,
		MapYLength: 103,
		Filename:   "inputs/input14.txt",
		Time:       100,
	}

	robotMap := day14

	data := getDay14Data(robotMap)
	easterEggData := getDay14Data(robotMap)

	robotPositions := cycleRobots(data, robotMap)
	safetyFactor := calculateSaferyFactor(robotPositions)
	log.Printf("The safty factor is %d", safetyFactor)

	findEasterEgg(&easterEggData, robotMap)
}

func findEasterEgg(robots *[]RobotGuard, robotMap RobotGuardMap) {
	for i := 0; i < 10000; i++ {
		results := make(map[Position]bool)
		for j := range *robots {
			robot := &(*robots)[j]
			robot.Move(robotMap)
			results[robot.Position] = true
		}
		printEasterEgg(results, robotMap, i+1)
	}
}

func printEasterEgg(robotMap map[Position]bool, robotGuardMap RobotGuardMap, index int) {
	possibleTree := false
	for y := 0; y < robotGuardMap.MapYLength; y++ {
		for x := 0; x < robotGuardMap.MapXLength; x++ {
			if checkIfPossibleTree(
				robotMap,
				y,
				x,
				robotGuardMap.MapYLength,
				robotGuardMap.MapXLength,
			) {
				possibleTree = true
			}
		}
	}

	if possibleTree {
		log.Printf("We found a possible tree at: %d", index)
		result := ""
		for y := 0; y < robotGuardMap.MapYLength; y++ {
			for x := 0; x < robotGuardMap.MapXLength; x++ {
				if robotMap[Position{y, x}] {
					result += "X"
				} else {
					result += " "
				}
			}
			result += "\n"
		}
		fmt.Println(result)
	}
}

func checkIfPossibleTree(robotMap map[Position]bool, row, col, rowLen, colLen int) bool {
	if row < 3 || row > rowLen-3 || col < 3 || col > colLen-3 {
		return false
	}

	for i := -2; i < 2; i++ {
		for j := -2; j < 2; j++ {
			if !robotMap[Position{row - i, col - j}] {
				return false
			}
		}
	}
	return true
}

func calculateSaferyFactor(robotMap map[string]int) int {
	result := 1
	for _, value := range robotMap {
		result *= value
	}
	return result
}

func cycleRobots(
	robots []RobotGuard,
	robotMap RobotGuardMap,
) map[string]int {
	results := make(map[string]int)
	yHalfway, xHalfway := robotMap.MapYLength/2, robotMap.MapXLength/2

	for _, robot := range robots {
		for i := 0; i < robotMap.Time; i++ {
			robot.Move(robotMap)
		}
		switch {
		case robot.Position.Row < yHalfway && robot.Position.Col < xHalfway:
			results["top-left"]++
		case robot.Position.Row < yHalfway && robot.Position.Col > xHalfway:
			results["top-right"]++
		case robot.Position.Row > yHalfway && robot.Position.Col < xHalfway:
			results["bottom-left"]++
		case robot.Position.Row > yHalfway && robot.Position.Col > xHalfway:
			results["bottom-right"]++
		}
	}

	return results
}

func (robot *RobotGuard) Move(robotMap RobotGuardMap) {
	newPosition := Position{
		robot.Position.Row + robot.Movement.Row,
		robot.Position.Col + robot.Movement.Col,
	}

	if newPosition.Row < 0 {
		newPosition.Row = newPosition.Row + robotMap.MapYLength
	}
	if newPosition.Col < 0 {
		newPosition.Col = newPosition.Col + robotMap.MapXLength
	}
	if newPosition.Row > robotMap.MapYLength-1 {
		newPosition.Row = newPosition.Row - robotMap.MapYLength
	}
	if newPosition.Col > robotMap.MapXLength-1 {
		newPosition.Col = newPosition.Col - robotMap.MapXLength
	}

	robot.Position = newPosition
}

func getDay14Data(robotGuardMap RobotGuardMap) []RobotGuard {
	file, err := os.Open(robotGuardMap.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var guards []RobotGuard
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		matches := robotGuardPattern.FindStringSubmatch(line)
		if len(matches) == 5 {
			pX, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Panicln("Error parsing position x", pX, err)
			}
			pY, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Panicln("Error parsing position y", pY, err)
			}
			vX, err := strconv.Atoi(matches[3])
			if err != nil {
				log.Panicln("Error parsing velocity x", vX, err)
			}
			vY, err := strconv.Atoi(matches[4])
			if err != nil {
				log.Panicln("Error parsing velocity y", vY, err)
			}

			guard := RobotGuard{
				// Do I want to keep this inverted?
				Position: Position{Row: pY, Col: pX},
				Movement: Position{Row: vY, Col: vX},
			}
			guards = append(guards, guard)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return guards
}
