package main

import (
	"aoc/days"
	"aoc/utils"
	"fmt"
	"log"
	"sync"
	"time"
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
		days.Day11,
		days.Day12,
		days.Day13,
		days.Day14,
		days.Day15,
		days.Day16,
		days.Day17,
		days.Day18,
		days.Day19,
		days.Day20,
		days.Day21,
		days.Day22,
		days.Day23,
		days.Day24,
		days.Day25,
	}

	var wg sync.WaitGroup

	today := time.Now().Day() - 4
	for i, dayFunc := range days {
		if i != today {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			dayFunc()
			timer.LogTime(fmt.Sprintf("Day %d", i+1))
		}()

		break
	}

	wg.Wait()

	timer.LogTotalTime()
}
