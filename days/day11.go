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
	plutoStones := viewPlutoStones(data)

	log.Println("stones", plutoStones)
}

func viewPlutoStones(data []int) int {
	results := make([]int, len(data))
	copy(results, data)

	for i := 0; i < 25; i++ {
		updatedResults := []int{}
		for j := 0; j < len(results); j++ {
			updatedResults = append(updatedResults, viewedStone(results[j])...)
		}
		results = updatedResults
	}
	return len(results)
}

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
