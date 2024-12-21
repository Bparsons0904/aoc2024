package days

import (
	"bufio"
	"log"
	"math"
	"os"
)

func Day20() {
	raceTrack := getDay20Data()
	getBaseRacePath(&raceTrack)
	shortestRoute := getShortestRoute(raceTrack)
	shortestRouteExpanded := getShortestRouteExpanded(raceTrack)
	log.Printf(
		"Number of routes that were reduced by atleast 100 picoseconds: %d and then those with crazy 20 picsecononds hack: %d",
		shortestRoute,
		shortestRouteExpanded,
	)
}

func getShortestRouteExpanded(raceTrack RaceTrack) int {
	distancesOver100 := 0

	for key, point := range raceTrack.Visited {
		visited := make(map[Position]bool)
		for row := -20; row <= 20; row++ {
			for col := -20; col <= 20; col++ {
				distanceApart := int(math.Abs(float64(row))) + int(math.Abs(float64(col)))
				if distanceApart > 20 {
					continue
				}
				direction := Position{row, col}
				if _, exists := visited[direction]; exists {
					continue
				}

				visited[direction] = true

				positionToCheck := key.GetNextMove(direction)
				if checkedValue, exists := raceTrack.Visited[positionToCheck]; exists {
					if checkedValue-distanceApart-point >= 100 {
						distancesOver100++
					}
				}
			}
		}
	}

	return distancesOver100
}

func getShortestRoute(raceTrack RaceTrack) int {
	distancesOver100 := 0

	for key, point := range raceTrack.Visited {
		for _, direction := range []Position{{0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			positionToCheck := key.GetNextMove(direction)
			if checkedValue, exists := raceTrack.Visited[positionToCheck]; exists {
				if checkedValue-2-point >= 100 {
					distancesOver100++
				}
			}

		}
	}

	return distancesOver100
}

func getBaseRacePath(raceTrack *RaceTrack) {
	i := 0
	for {
		nextPosition := raceTrack.getNextPosition()
		raceTrack.Visited[nextPosition] = raceTrack.Visited[raceTrack.Position] + 1
		raceTrack.Position = nextPosition

		if raceTrack.Map[nextPosition] == end {
			break
		}

		i++
	}
}

func (raceTrack *RaceTrack) getNextPosition() Position {
	for _, direction := range Directions[:4] {
		position := raceTrack.Position.GetNextPosition(direction)
		_, visited := raceTrack.Visited[position]
		toCheck := raceTrack.Map[position]
		if toCheck == space && !visited || toCheck == end {
			return position
		}
	}

	return Position{}
}

type RaceTrack struct {
	Map      PosMap
	Visited  map[Position]int
	RowLen   int
	ColLen   int
	Position Position
}

func getDay20Data() RaceTrack {
	file, err := os.Open("inputs/input20.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	raceTrack := RaceTrack{
		Map:     make(PosMap),
		Visited: make(map[Position]int),
		RowLen:  0,
		ColLen:  0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()

		for j, val := range row {
			raceTrack.Map[Position{raceTrack.RowLen, j}] = val
			if val == start {
				raceTrack.Position = Position{raceTrack.RowLen, j}
				raceTrack.Visited[raceTrack.Position] = 0
			}
		}

		if raceTrack.RowLen == 0 {
			raceTrack.ColLen = len(row)
		}
		raceTrack.RowLen++
	}

	return raceTrack
}
