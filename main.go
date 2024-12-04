package main

import (
	"aoc/utils"
	"log"
	"sync"
)

func main() {
	log.Println("Hello AoC")
	timer := utils.StartTimer("Aoc")

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		day1()
		timer.LogTime("Day 1")
	}()

	go func() {
		defer wg.Done()
		day2()
		timer.LogTime("Day 2")
	}()

	go func() {
		defer wg.Done()
		day3()
		timer.LogTime("Day 3")
	}()

	wg.Wait()

	timer.LogTotalTime()
}
