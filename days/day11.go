package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	data := getDay11Data()
	stoneCounts := make(map[int]int)
	for _, stone := range data {
		stoneCounts[stone]++
	}

	blinks25 := viewPlutoStonesWithMap(stoneCounts, 25)
	blinks75 := viewPlutoStonesWithMap(stoneCounts, 75)
	log.Printf(
		"Total stones after 25 blinks: %d and %d stones with 75 freaking blinks",
		getTotalStones(blinks25),
		getTotalStones(blinks75),
	)
}

func viewPlutoStonesWithMap(stoneCounts map[int]int, views int) map[int]int {
	currentCounts := make(map[int]int)
	for key, value := range stoneCounts {
		currentCounts[key] = value
	}

	for i := 0; i < views; i++ {
		newCounts := make(map[int]int)
		for stone, count := range currentCounts {
			newStones := viewedStone(stone)
			for _, newStone := range newStones {
				newCounts[newStone] += count
			}
		}
		currentCounts = newCounts
	}
	return currentCounts
}

func getTotalStones(stoneCounts map[int]int) int {
	total := 0
	for _, count := range stoneCounts {
		total += count
	}
	return total
}

// func viewPlutoStonesTimes3(plutoStones []int) int {
// 	finalResults := 0
//
// 	for _, stone := range plutoStones {
// 		log.Println("stone", stone)
// 		finalResults += viewStoneRecursively(stone, 75, 0)
// 	}
// 	// for i := 0; i < len(plutoStones); i++ {
// 	// 	group := viewPlutoStones(plutoStones[i:i+1], 25)
// 	// 	for j := 0; j < len(group); j++ {
// 	// 		log.Println("finalResults", j, len(group), len(plutoStones))
// 	// 		finalResults += len(viewPlutoStones(group[j:j+1], 25))
// 	// 	}
// 	// 	log.Println("group", i, len(group), len(plutoStones))
// 	// }
// 	return finalResults
// }
//
// func viewStoneRecursively(stone, iteration, result int) int {
// 	if iteration == 0 {
// 		return 1
// 	}
//
// 	results := 0
// 	viewedStones := viewedStone(stone)
// 	log.Println("viewedStones", viewedStones, iteration, result)
// 	for _, viewStone := range viewedStones {
// 		results += viewStoneRecursively(viewStone, iteration-1, result)
// 	}
//
// 	return results
// }
//
// func viewPlutoStones(data []int, views int) []int {
// 	// log.Println("Starting view", len(data), views)
// 	results := make([]int, len(data))
// 	copy(results, data)
//
// 	for i := 0; i < views; i++ {
// 		updatedResults := []int{}
// 		// log.Printf("Working on round %d, with length of %d", i, len(results))
// 		// log.Println("Actual array", i, results)
// 		for j := 0; j < len(results); j++ {
// 			updatedResults = append(updatedResults, viewedStone(results[j])...)
// 		}
// 		results = updatedResults
// 	}
// 	return results
// }

func viewedStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	stoneString := strconv.Itoa(stone)
	if len(stoneString)%2 == 0 {
		first := stoneString[0 : len(stoneString)/2]
		second := stoneString[(len(stoneString) / 2):]
		firstInt, err := strconv.Atoi(first)
		if err != nil {
			log.Panicln("Error converting first to int", stone, first)
		}
		secondInt, err := strconv.Atoi(second)
		if err != nil {
			log.Panicln("Error converting second to int", stone, second)
		}
		return []int{firstInt, secondInt}
	}

	return []int{stone * 2024}
}

func getDay11Data() []int {
	file, err := os.Open("inputs/input11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stones := strings.Split(scanner.Text(), " ")

		for _, stone := range stones {
			num, err := strconv.Atoi(string(stone))
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, num)
		}

	}

	return results
}
