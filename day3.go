package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var (
	mulRegex    = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	mulDoRegex  = regexp.MustCompile(`(?s)don't\(\).*?(?:do\(\)|$)`)
	numberRegex = regexp.MustCompile(`\d{1,3}`)
)

func day3() {
	data := getDay3Data()

	var wg sync.WaitGroup
	var rentalCalulations, rentalValidatedCalculations int
	var mulDoString string

	wg.Add(1)
	go func() {
		defer wg.Done()
		mulString := mulRegex.FindAllStringSubmatch(data, -1)
		pairData := getAllPairs(mulString)
		rentalCalulations = processMulCalculations(pairData)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mulDoString = mulDoRegex.ReplaceAllLiteralString(data, "")
		mulDoStringResults := mulRegex.FindAllStringSubmatch(mulDoString, -1)
		pairData := getAllPairs(mulDoStringResults)
		rentalValidatedCalculations = processMulCalculations(pairData)
	}()

	wg.Wait()

	log.Printf(
		"Rental Calculations: %d\nValidated Calculations: %d",
		rentalCalulations,
		rentalValidatedCalculations,
	)
}

func processMulCalculations(data []int) int {
	total := 0
	for i := 0; i < len(data); i += 2 {
		total += data[i] * data[i+1]
	}

	return total
}

func getDay3Data() string {
	file, err := os.ReadFile("inputs/input3.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}

func getAllPairs(testString [][]string) []int {
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
