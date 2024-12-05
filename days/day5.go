package days

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	correctedInvalidTotals, correctlyOrderedMidTotal := 0, 0
	rulesOrderingMap, pageNumbers := getDay5Data()

PageLoop:
	for _, page := range pageNumbers {
		for i, number := range page {
			if orderingMap, ok := rulesOrderingMap[number]; ok {
				if valid, _ := validateOrder(page[:i], orderingMap); !valid {
					correctedOrder := fixInvalidOrders(page, rulesOrderingMap)
					correctedInvalidTotals += correctedOrder[len(correctedOrder)/2]
					continue PageLoop
				}
			}
		}

		correctlyOrderedMidTotal += page[len(page)/2]
	}

	log.Printf(
		"Found Validate Ordered Printing Total: %d and Corrected Ordered Total: %d",
		correctlyOrderedMidTotal,
		correctedInvalidTotals,
	)
}

func fixInvalidOrders(invalidPage []int, rulesOrderingMap map[int][]int) []int {
	for i, number := range invalidPage {
		if orderingMap, ok := rulesOrderingMap[number]; ok {
			valid, outOfOrder := validateOrder(invalidPage[:i], orderingMap)
			if !valid {
				invalidPage[i], invalidPage[outOfOrder] = invalidPage[outOfOrder], invalidPage[i]
				return fixInvalidOrders(invalidPage, rulesOrderingMap)
			}

		}
	}

	return invalidPage
}

func validateOrder(remainingNumbers []int, orderingMap []int) (bool, int) {
	for i, number := range remainingNumbers {
		if slices.Contains(orderingMap, number) {
			return false, i
		}
	}

	return true, 0
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

	return rulesOrderingMap, printList
}
