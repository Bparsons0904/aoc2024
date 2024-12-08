package days

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
)

var nodeRegex = regexp.MustCompile(`\w`)

func Day8() {
	data, mapRowLen, mapColLen := getDay8Data()
	uniqueNodes := []Direction{}

	for _, nodes := range data {
		for i, node := range nodes {
			for j := i; j < len(nodes)-1; j++ {
				antiNodes := nodeToCheck(node, nodes[j+1], mapRowLen, mapColLen)
				for _, antiNode := range antiNodes {
					if !slices.ContainsFunc(uniqueNodes, func(dir Direction) bool {
						return dir.Row == antiNode.Row && dir.Col == antiNode.Col
					}) {
						uniqueNodes = append(uniqueNodes, antiNode)
					}
				}
			}
		}
	}

	log.Println(len(uniqueNodes))
}

func filterDuplicateDirections(directions []Direction) []Direction {
	results := []Direction{}
	for _, direction := range directions {
		if !slices.ContainsFunc(results, func(dir Direction) bool {
			return dir.Row == direction.Row && dir.Col == direction.Col
		}) {
			results = append(results, direction)
		}
	}
	return results
}

func nodeToCheck(nodeA, nodeB Direction, rowLen, colLen int) []Direction {
	dif := Direction{nodeA.Row - nodeB.Row, nodeA.Col - nodeB.Col}
	toCheck := []Direction{
		{nodeA.Row + dif.Row, nodeA.Col + dif.Col},
		{nodeB.Row - dif.Row, nodeB.Col - dif.Col},
	}
	results := []Direction{}
	for _, node := range toCheck {
		if isInBounds(node.Row, node.Col, rowLen, colLen) {
			results = append(results, node)
		}
	}
	return results
}

func getDay8Data() (map[byte][]Direction, int, int) {
	file, err := os.Open("inputs/input8.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	results := make(map[byte][]Direction)
	i := 0
	mapColLen := 0
	for scanner.Scan() {
		row := scanner.Text()
		if i == 0 {
			mapColLen = len(row)
		}
		for j := 0; j < len(row); j++ {
			if nodeRegex.Match([]byte{row[j]}) {
				results[row[j]] = append(results[row[j]], Direction{i, j})
			}
		}
		i++
	}

	return results, i, mapColLen
}
