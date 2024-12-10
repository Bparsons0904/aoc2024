package main

import (
	"aoc/days"
	"aoc/utils"
	"log"
	"sync"
)

func main() {
	log.Println("Hello AoC")
	timer := utils.StartTimer("Aoc")

	var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day1()
	// 	timer.LogTime("Day 1")
	// }()
	//
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day2()
	// 	timer.LogTime("Day 2")
	// }()
	//
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day3()
	// 	timer.LogTime("Day 3")
	// }()
	//
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day4()
	// 	timer.LogTime("Day 4")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day5()
	// 	timer.LogTime("Day 5")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day6()
	// 	timer.LogTime("Day 6")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day7()
	// 	timer.LogTime("Day 7")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	days.Day8()
	// 	timer.LogTime("Day 8")
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		days.Day9()
		timer.LogTime("Day 9")
	}()

	wg.Wait()

	timer.LogTotalTime()
}
