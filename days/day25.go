package days

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day25() {
	data := getDay25Data()
	log.Println(data)
}

func getDay25Data() []int {
	file, err := os.Open("inputs/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stones := strings.Split(scanner.Text(), " ")

		for _, stone := range stones {
			num, err := strconv.Atoi(string(stone))
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, num)
		}

	}

	return results
}
