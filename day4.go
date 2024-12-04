package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func day4() {
	_ = getDay4Data()

	log.Printf(
		"Rental Calculations: %d\nValidated Calculations: %d",
		0, 1,
	)
}

func getDay4Data() string {
	file, err := os.Open("inputs/input4.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		log.Panicln("row", row)
	}

	return "test"
}
