package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Data   [][]int
	Safe   int
	UnSafe int
}

func day2() {
	var report Report
	getData(&report)

	calculateReportSafety(&report)
	log.Println("Safe Reports", report.Safe)
}

func calculateReportSafety(report *Report) {
	for _, row := range report.Data {
		test := row[1] - row[0]
		switch {
		case test == 0:
			report.UnSafe++
		case test > 0:
			calculateIncreasingReport(report, row)
		case test < 0:
			calculateDecreasingReport(report, row)
		}
	}
}

func calculateIncreasingReport(report *Report, row []int) {
	log.Println("Increasing test", row)
	for i := 0; i < len(row)-1; i++ {
		difference := row[i+1] - row[i]
		if difference <= 0 || difference > 3 {
			report.UnSafe++
			return
		}
	}

	report.Safe++
}

func calculateDecreasingReport(report *Report, row []int) {
	for i := 0; i < len(row)-1; i++ {
		difference := row[i] - row[i+1]
		if difference <= 0 || difference > 3 {
			report.UnSafe++
			return
		}
	}

	report.Safe++
}

func getData(report *Report) {
	file, err := os.Open("inputs/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		intRow := make([]int, len(row))

		for i, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			intRow[i] = num
		}

		report.Data = append(report.Data, intRow)
	}
}
