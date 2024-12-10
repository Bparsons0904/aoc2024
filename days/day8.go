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
	uniqueNodesExtended := []Direction{}

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
				antiNodesExtendedToAdd := nodeToCheckExtended(
					node,
					nodes[j+1],
					mapRowLen,
					mapColLen,
				)
				for _, antiNode := range antiNodesExtendedToAdd {
					if !slices.ContainsFunc(uniqueNodesExtended, func(dir Direction) bool {
						return dir.Row == antiNode.Row && dir.Col == antiNode.Col
					}) {
						uniqueNodesExtended = append(uniqueNodesExtended, antiNode)
					}
				}
			}
		}
	}

	log.Println(len(uniqueNodes))
	log.Println(len(uniqueNodesExtended))
}

func nodeToCheckExtended(nodeA, nodeB Direction, rowLen, colLen int) []Direction {
	dif := Direction{nodeA.Row - nodeB.Row, nodeA.Col - nodeB.Col}
	results := []Direction{}

	i := 1
	for {
		toCheck := Direction{
			nodeA.Row + (i * dif.Row), nodeA.Col + i*(dif.Col),
		}
		if isInBounds(toCheck.Row, toCheck.Col, rowLen, colLen) {
			results = append(results, toCheck)
			i++
			continue
		}
		break
	}
	j := 1

	for {
		toCheck := Direction{
			nodeB.Row - (j * dif.Row), nodeB.Col - j*(dif.Col),
		}
		if isInBounds(toCheck.Row, toCheck.Col, rowLen, colLen) {
			results = append(results, toCheck)
			j++
			continue
		}
		break
	}

	results = append(results, nodeA)
	results = append(results, nodeB)
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
