package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var correctlyOrderedMidTotal = 0

func Day5() {
	rulesOrderingMap, pageNumbers := getDay5Data()
	log.Println(rulesOrderingMap, pageNumbers)

PageLoop:
	for _, page := range pageNumbers {
		log.Println("Starting page loop", page)
		for i, number := range page {

			log.Println("Starting number loop", number)
			if orderingMap, ok := rulesOrderingMap[number]; ok {
				if !validateOrder(page[:i], orderingMap) {
					log.Println("We have a invalide ordeer")
					continue PageLoop
				}
			}
		}

		correctlyOrderedMidTotal += page[len(page)/2]
	}

	log.Printf("Found Validate Ordered Printing Total: %d", correctlyOrderedMidTotal)
}

func validateOrder(remainingNumbers []int, orderingMap []int) bool {
	// log.Println("validate", remainingNumbers)
	for _, number := range remainingNumbers {
		if slices.Contains(orderingMap, number) {
			return false
		}
	}

	return true
}

func getDay5Data() (map[int][]int, [][]int) {
	file, err := os.Open("inputs/input5.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	rulesOrderingMap := make(map[int][]int)
	printList := [][]int{}
	generateMap := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			generateMap = false
			continue
		}

		addToMap := func(line string) {
			rules := strings.Split(line, "|")
			log.Println("rules", rules)

			beforeRule, err := strconv.Atoi(rules[0])
			if err != nil {
				log.Panicln("Error parsing before rule", err)
			}
			afterRule, err := strconv.Atoi(rules[1])
			if err != nil {
				log.Panicln("Error parsing after rule", err)
			}

			rulesOrderingMap[beforeRule] = append(rulesOrderingMap[beforeRule], afterRule)
		}

		addToList := func(line string) {
			list := strings.Split(line, ",")
			intList := []int{}
			for _, pageString := range list {
				pageNumber, err := strconv.Atoi(pageString)
				if err != nil {
					log.Panicln("Error parsing page number", err)
				}

				intList = append(intList, pageNumber)
			}

			printList = append(printList, intList)
		}

		if generateMap {
			addToMap(line)
		} else {
			addToList(line)
		}

	}

	log.Println("rule", rulesOrderingMap, printList)

	return rulesOrderingMap, printList
}
