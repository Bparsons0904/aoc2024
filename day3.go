package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func day3() {
	data, dataExpanded := getDay3Data()

	allValues := processMulCalculations(data)
	dontValues := processMulCalculations(dataExpanded)

	log.Printf(
		"Total Multiplications: %d\nTotal Don't Values: %d",
		allValues,
		dontValues,
	)
}

func processMulCalculations(data []int) int {
	total := 0
	for i := 0; i < len(data); i += 2 {
		total += data[i] * data[i+1]
	}

	return total
}

func getDay3Data() ([]int, []int) {
	file, err := os.ReadFile("inputs/input3.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(file)
	mulRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	mulString := mulRegex.FindAllStringSubmatch(fileString, -1)
	result := getAllPairs(mulString)

	mulDoRegex := regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)
	mulDoString := mulDoRegex.ReplaceAllLiteralString(fileString, "")
	mulDoStringResults := mulRegex.FindAllStringSubmatch(mulDoString, -1)
	result2 := getAllPairs(mulDoStringResults)

	return result, result2
}

func getAllPairs(testString [][]string) []int {
	numberRegex := regexp.MustCompile(`\d{1,3}`)
	results := []int{}
	for _, pair := range testString {
		intRegex := numberRegex.FindAllStringSubmatch(pair[0], -1)

		first, err := strconv.Atoi(intRegex[0][0])
		if err != nil {
			log.Panicln("Failed to parse first int", pair)
		}

		second, err := strconv.Atoi(intRegex[1][0])
		if err != nil {
			log.Panicln("Failed to parse second int", pair)
		}
		results = append(results, first)
		results = append(results, second)
	}

	return results
}
