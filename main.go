package main

import (
	"aoc/days"
	"aoc/utils"
	"fmt"
	"log"
	"sync"
)

type DayFunc func()

func main() {
	log.Println("Hello AoC")
	timer := utils.StartTimer("Aoc")
	days := []DayFunc{
		days.Day1,
		days.Day2,
		days.Day3,
		days.Day4,
		days.Day5,
		days.Day6,
		days.Day7,
		days.Day8,
		days.Day9,
		days.Day10,
	}

	var wg sync.WaitGroup

	startPoint := len(days) - 1
	for i, dayFunc := range days {
		if i < startPoint {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			dayFunc()
			timer.LogTime(fmt.Sprintf("Day %d", i+1))
		}()

	}

	wg.Wait()

	timer.LogTotalTime()
}
