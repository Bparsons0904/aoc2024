package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

)

func day3() {
	data := getDay3Data()
}

func getDay3Data() []int {
	file, err := os.Open("inputs/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		regex := "mul\(d e{1-3},d{1-3}\)"
		intRow := make([]int, len(row))

		for i, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			intRow[i] = num
		}

	}
}
