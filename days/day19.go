package days

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day19() {
	designs, patterns := getDay19Data()
	log.Println(designs, patterns)
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
