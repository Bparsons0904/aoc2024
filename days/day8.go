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
	// log.Println("data", data, mapRowLen, mapColLen)
	// nodeMap := make(map[byte][]Direction)
	uniqueNodes := []Direction{}

	for _, nodes := range data {
		// log.Println("uniqueNodes", uniqueNodes)
		// log.Println("nodes", nodes)
		for i, node := range nodes {
			for j := i; j < len(nodes)-1; j++ {
				// for j := i; j < len(nodes)-1; j++ {
				// log.Println(key, j, nodes[j])
				antiNodes := nodeToCheck(node, nodes[j+1], mapRowLen, mapColLen)
				// nodeMap[key] = append(nodeMap[key], antiNodes...)
				// log.Println("We are going to check uniqueNodes", antiNodes)
				for _, antiNode := range antiNodes {
					if !slices.ContainsFunc(uniqueNodes, func(dir Direction) bool {
						return dir.Row == antiNode.Row && dir.Col == antiNode.Col
					}) {
						// log.Println("Adding these nodes", antiNode)
						uniqueNodes = append(uniqueNodes, antiNode)
					}
					// nodeMap[key] = append(nodeMap[key], antiNode)
				}
				// }
			}
		}
	}

	// totalAntiNode := 0
	// for _, nodes := range nodeMap {
	// 	// log.Println("nodes", nodes)
	// 	filteredNodes := filterDuplicateDirections(nodes)
	// 	// log.Println("filteredNodes", filteredNodes)
	// 	totalAntiNode += len(filteredNodes)
	// }

	// log.Println(uniqueNodes)
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
	// log.Println("To check original nodes", nodeA, nodeB)
	dif := Direction{nodeA.Row - nodeB.Row, nodeA.Col - nodeB.Col}
	// log.Println("dif", dif)

	toCheck := []Direction{
		{nodeA.Row + dif.Row, nodeA.Col + dif.Col},
		{nodeB.Row - dif.Row, nodeB.Col - dif.Col},
	}
	// if dif.Col < 0 {
	// } else {
	// 	toCheck = []Direction{
	// 		{nodeA.Row + dif.Row, nodeA.Col + dif.Col},
	// 		{nodeB.Row - dif.Row, nodeB.Col - dif.Col},
	// 	}
	// }
	// log.Println("toCheck", toCheck)
	results := []Direction{}
	for _, node := range toCheck {
		// log.Println("Checking in bounds", node)
		if isInBounds(node.Row, node.Col, rowLen, colLen) {
			results = append(results, node)
		}
	}
	// log.Println("results", results)
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
