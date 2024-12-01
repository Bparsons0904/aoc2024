package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Places struct {
	AID []float64
	BID []float64
}

func day1() {
	places := getPlaces()
	day1a(places)
	day1b(places)
}

func day1a(places Places) {
	difference := 0.0
	for i := 0; i < len(places.AID); i++ {
		a := places.AID[i]
		b := places.BID[i]

		difference += math.Abs(a - b)
	}

	log.Printf(
		"Day 1A, Total difference: %d",
		int(difference),
	)
}

func day1b(places Places) {
	placeMap := make(map[float64]int)

	for _, aID := range places.AID {
		placeMap[aID] = 0
	}

	for _, bID := range places.BID {
		_, ok := placeMap[bID]
		if ok {
			placeMap[bID] += 1
		}
	}

	total := 0
	for key, value := range placeMap {
		total += int(key) * value
	}

	log.Printf(
		"Day 1B, Total: %d",
		total,
	)
}

func getPlaces() Places {
	file, err := os.Open("inputs/input1a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var places Places
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		a, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		places.AID = append(places.AID, a)
		places.BID = append(places.BID, b)
	}

	slices.Sort(places.AID)
	slices.Sort(places.BID)

	return places
}
