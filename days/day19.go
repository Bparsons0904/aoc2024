package days

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day19() {
	designs, patterns := getDay19Data()
	// log.Println(hasPattern(designs[0], patterns))
	// log.Println(calculateDesignsWithPatterns(designs[:3], patterns))
	log.Println(calculateDesignsWithPatterns(designs, patterns))
	// log.Println(patterns)
}

// func calculateTotalDesignsWithPatterns(designs, patterns []string) int {
// 	result := 0
// 	for _, design := range designs {
// 		result += hasUniquePattern(design, patterns)
// 	}
//
// 	return result
// }
//
// func hasUniquePattern(design string, patterns []string, currentPattern []string, existingFoundPatterns [][]string) []string {
// 	if design == "" {
// return currentPattern
// 	}
//
// 	for _, pattern := range patterns {
// 		if strings.HasPrefix(design, pattern) {
// 			newDesign := design[len(pattern):]
// 			if hasPattern(newDesign, patterns) {
// 				return append()
// 			}
// 		}
// 	}
//
//
// 	return []string{}
// }

func calculateDesignsWithPatterns(designs, patterns []string) int {
	result := 0
	for _, design := range designs {
		if hasPattern(design, patterns) {
			result++
		}
	}

	return result
}

func hasPattern(design string, patterns []string) bool {
	if design == "" {
		return true
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			newDesign := design[len(pattern):]
			if hasPattern(newDesign, patterns) {
				return true
			}
		}
	}
	return false
}

func getDay19Data() ([]string, []string) {
	file, err := os.Open("inputs/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	patterns := []string{}
	designs := []string{}
	scanner := bufio.NewScanner(file)
	scanningDesigns := false

	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			scanningDesigns = true
			continue
		}

		if !scanningDesigns {
			patterns = strings.Split(row, ", ")
			continue
		}

		designs = append(designs, row)

	}

	return designs, patterns
}
