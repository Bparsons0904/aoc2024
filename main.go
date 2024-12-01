package main

import (
	"aoc/utils"
	"log"
)

func main() {
	log.Println("Hello AoC")
	timer := utils.StartTimer("Aoc")

	day1()
	timer.LogTime("Day 1")

	timer.LogTotalTime()
}
