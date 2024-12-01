package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func day2a() {
	start := time.Now()
	places := getPlaces()

	placeMap := make(map[float64]int)

	for _, aID := range places.AID {
		for _, bID := range places.BID {
			if aID == bID {
				placeMap[aID] += 1
			}
		}
	}

	total := 0
	for key, value := range placeMap {
		total += int(key) * value
	}

	fmt.Printf(
		"Day 1, problem B took: %v\nTotal: %d",
		time.Since(start),
		total,
	)
}

func day1a() {
	start := time.Now()
	places := getPlaces()

	difference := 0.0
	for i := 0; i < len(places.AID); i++ {
		a := places.AID[i]
		b := places.BID[i]

		difference += math.Abs(a - b)
	}

	fmt.Printf(
		"Day 1, problem A took: %v\nTotal difference: %d",
		time.Since(start),
		int(difference),
	)
}

type Places struct {
	AID []float64
	BID []float64
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
		line := scanner.Text()
		fields := strings.Split(line, "   ")
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
