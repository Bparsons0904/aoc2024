package days

import (
	"bufio"
	"log"
	"os"
)

func Day12() {
	data := getDay12Data()
	plotMap := getPlotMap(data)
	fencePricing := calculateFencePricing(plotMap)
	fencePricingWithDiscount := calculateFencePricingWithDiscount(plotMap)

	log.Printf(
		"Total cost for fencing for aall regions %d with the discounted price of %d",
		fencePricing,
		fencePricingWithDiscount,
	)
}

func calculateFencePricingWithDiscount(plotMap map[rune][][]Position) int {
	totals := 0
	for _, plotAreas := range plotMap {
		for _, plotArea := range plotAreas {
			perimeter := 0
			area := len(plotArea)
			for _, plot := range plotArea {
				plotPerimeter := getPerimeter(plot, plotArea)
				perimeter += plotPerimeter
			}
			totals += area * perimeter
		}
	}

	return totals
}

func calculateFencePricing(plotMap map[rune][][]Position) int {
	totals := 0

	for _, plotAreas := range plotMap {
		for _, plotArea := range plotAreas {
			perimeter := 0
			for _, plot := range plotArea {
				touches := plot.TouchesAnotherPosition(plotArea)
				perimeter += 4 - touches

			}
			area := len(plotArea)
			totals += area * perimeter
		}
	}

	return totals
}

func getPerimeter(plot Position, plotArea []Position) int {
	perimeter := 0

	isTopEdge := true
	for _, test := range plotArea {
		if plot.Col == test.Col && plot.Row-1 == test.Row {
			isTopEdge = false
		}
	}

	if isTopEdge {
		for _, test := range plotArea {
			if plot.Col == test.Col+1 && plot.Row == test.Row {
				isTopAlso := false
				for _, test2 := range plotArea {
					if test.Col == test2.Col && test.Row-1 == test2.Row {
						isTopAlso = true
						break
					}
				}

				if !isTopAlso {
					isTopEdge = false
				}
				break
			}
		}
	}

	if isTopEdge {
		perimeter++
	}

	isBottomEdge := true
	for _, test := range plotArea {
		if plot.Col == test.Col && plot.Row == test.Row-1 {
			isBottomEdge = false
		}
	}

	if isBottomEdge {
		for _, test := range plotArea {
			if plot.Col == test.Col+1 && plot.Row == test.Row {
				isBottomAlso := false
				for _, test2 := range plotArea {
					if test.Col == test2.Col && test.Row == test2.Row-1 {
						isBottomAlso = true
						break
					}
				}

				if !isBottomAlso {
					isBottomEdge = false
				}
				break
			}
		}
	}

	if isBottomEdge {
		perimeter++
	}

	isLeftEdge := true
	for _, test := range plotArea {
		if plot.Row == test.Row && plot.Col-1 == test.Col {
			isLeftEdge = false
		}
	}

	if isLeftEdge {
		for _, test := range plotArea {
			if plot.Col == test.Col && plot.Row == test.Row+1 {
				isLeftAlso := false
				for _, test2 := range plotArea {
					if test.Col-1 == test2.Col && test.Row == test2.Row {
						isLeftAlso = true
						break
					}
				}

				if !isLeftAlso {
					isLeftEdge = false
				}
				break
			}
		}
	}

	if isLeftEdge {
		perimeter++
	}

	isRightEdge := true
	for _, test := range plotArea {
		if plot.Row == test.Row && plot.Col+1 == test.Col {
			isRightEdge = false
		}
	}

	if isRightEdge {
		for _, test := range plotArea {
			if plot.Col == test.Col && plot.Row == test.Row+1 {
				isRightAlso := false
				for _, test2 := range plotArea {
					if test.Col+1 == test2.Col && test.Row == test2.Row {
						isRightAlso = true
						break
					}
				}

				if !isRightAlso {
					isRightEdge = false
				}
				break
			}
		}
	}

	if isRightEdge {
		perimeter++
	}

	return perimeter
}

func getPlotMap(data []string) map[rune][][]Position {
	plotMap := make(map[rune][][]Position)
	visited := make(map[Position]bool)

	for i := 0; i < len(data); i++ {
		for j, plot := range data[i] {
			pos := Position{i, j}
			if !visited[pos] {
				group := findTouchingPlots(pos, plot, visited, data)
				if len(group) > 0 {
					plotMap[plot] = append(plotMap[plot], group)
				}
			}
		}
	}

	return plotMap
}

func findTouchingPlots(
	position Position,
	plot rune,
	visited map[Position]bool,
	data []string,
) []Position {
	if visited[position] || rune(data[position.Row][position.Col]) != plot {
		return nil
	}

	group := []Position{position}
	visited[position] = true

	for _, dir := range directions[:4] {
		nextPos := Position{position.Row + dir.Row, position.Col + dir.Col}
		if !checkPositionInBounds(nextPos, Position{len(data), len(data[0])}) {
			continue
		}
		connected := findTouchingPlots(nextPos, plot, visited, data)
		group = append(group, connected...)
	}

	return group
}

func (position Position) TouchesAnotherPosition(positions []Position) int {
	touches := 0
	for _, test := range positions {
		for _, direction := range directions[:4] {

			positionCheck := test.Plus(direction)
			if positionCheck == position {
				touches++
			}
		}
	}
	return touches
}

func (position Position) Plus(direction Position) Position {
	return Position{Row: position.Row + direction.Row, Col: position.Col + direction.Col}
}

func getDay12Data() []string {
	file, err := os.Open("inputs/input12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		plots := scanner.Text()
		results = append(results, plots)
	}

	return results
}
