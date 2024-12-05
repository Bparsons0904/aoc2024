package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Data         [][]int
	Safe         int
	DampenedSafe int
}

func Day2() {
	var report Report
	getData(&report)

	calculateReportSafety(&report)
	log.Printf("\nSafe Reports: %d\nSafe w/ Dampened: %d", report.Safe, report.DampenedSafe)
}

func calculateReportSafety(report *Report) {
	for _, row := range report.Data {
		test := row[1] - row[0]
		switch {
		case test > 0:
			if testIncreasing(row) {
				report.Safe++
				report.DampenedSafe++
			} else {
				calculateReportWithDampener(report, row)
			}
		case test <= 0:
			if testDecreasing(row) {
				report.Safe++
				report.DampenedSafe++
			} else {
				calculateReportWithDampener(report, row)
			}
		}
	}
}

func calculateReportWithDampener(report *Report, row []int) {
	for i := 0; i < len(row); i++ {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		updatedRow := append(rowCopy[:i], rowCopy[i+1:]...)
		test := updatedRow[1] - updatedRow[0]
		switch {
		case test > 0:
			if testIncreasing(updatedRow) {
				report.DampenedSafe++
				return
			}
		case test < 0:
			if testDecreasing(updatedRow) {
				report.DampenedSafe++
				return
			}
		}
	}
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

func testIncreasing(row []int) bool {
	if ((row[0]-row[len(row)-1])*1)/len(row) > 3 {
		return false
	}
	for i := 0; i < len(row)-1; i++ {
		difference := row[i+1] - row[i]
		if difference <= 0 || difference > 3 {
			return false
		}
	}

	return true
}

func testDecreasing(row []int) bool {
	if ((row[0]-row[len(row)-1])*1)/len(row) > 3 {
		return false
	}
	for i := 0; i < len(row)-1; i++ {
		difference := row[i] - row[i+1]
		if difference <= 0 || difference > 3 {
			return false
		}
	}

	return true
}
