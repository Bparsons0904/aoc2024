package days

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
)

func Day7() {
	data := getDay7Data()
	calibrationTotals := make(chan *big.Int, len(data))
	calibrationsToRetotal := make(chan []*big.Int, len(data))
	var wg sync.WaitGroup

	for _, calibrations := range data {
		wg.Add(1)
		go func(calibs []*big.Int) {
			defer wg.Done()
			if result := processCalibration(calibs); result != nil {
				calibrationTotals <- result
			} else {
				calibrationsToRetotal <- calibs
			}
		}(calibrations)
	}

	wg.Wait()
	close(calibrationsToRetotal)

	var retotalSlice [][]*big.Int
	for calibs := range calibrationsToRetotal {
		retotalSlice = append(retotalSlice, calibs)
	}

	expandedTotals := make(chan *big.Int, len(retotalSlice))
	for _, calibrations := range retotalSlice {
		wg.Add(1)
		go func(calibs []*big.Int) {
			defer wg.Done()
			if result := procExpandedCalculations(calibs); result != nil {
				expandedTotals <- result
			}
		}(calibrations)
	}

	wg.Wait()
	close(calibrationTotals)
	close(expandedTotals)

	totalSum := big.NewInt(0)
	expandedSum := big.NewInt(0)

	for result := range calibrationTotals {
		totalSum.Add(totalSum, result)
	}

	for result := range expandedTotals {
		expandedSum.Add(expandedSum, result)
	}

	finalSum := new(big.Int).Add(totalSum, expandedSum)
	log.Printf(
		"Total of calibrations: %s with an expanded Calibration Total of: %s",
		totalSum.String(),
		finalSum.String(),
	)
}

func processCalibration(calibrations []*big.Int) *big.Int {
	expectedTotal := calibrations[0]
	potentials := []*big.Int{calibrations[1]}

	for i := 2; i < len(calibrations); i++ {
		updatedPotentials := []*big.Int{}
		for _, potential := range potentials {
			add := new(big.Int).Add(calibrations[i], potential)
			if add.Cmp(expectedTotal) <= 0 {
				updatedPotentials = append(updatedPotentials, add)
			}
			multiply := new(big.Int).Mul(calibrations[i], potential)
			if multiply.Cmp(expectedTotal) <= 0 {
				updatedPotentials = append(updatedPotentials, multiply)
			}
			potentials = updatedPotentials
		}
	}

	for _, result := range potentials {
		if result.Cmp(expectedTotal) == 0 {
			return result
		}
	}
	return nil
}

func procExpandedCalculations(calibrations []*big.Int) *big.Int {
	expectedTotal := calibrations[0]
	potentials := []*big.Int{calibrations[1]}

	for i := 2; i < len(calibrations); i++ {
		updatedPotentials := []*big.Int{}
		addToPotentials := func(value *big.Int) {
			if value.Cmp(expectedTotal) <= 0 {
				updatedPotentials = append(updatedPotentials, value)
			}
		}

		for _, potential := range potentials {
			add := new(big.Int).Add(calibrations[i], potential)
			addToPotentials(add)
			multiply := new(big.Int).Mul(calibrations[i], potential)
			addToPotentials(multiply)
			concat := concatInts(potential, calibrations[i])
			addToPotentials(concat)
			addConcat := concatInts(potential, add)
			addToPotentials(addConcat)
			mulConcat := concatInts(potential, multiply)
			addToPotentials(mulConcat)
			potentials = updatedPotentials
		}
	}

	for _, result := range potentials {
		if result.Cmp(expectedTotal) == 0 {
			return result
		}
	}
	return nil
}

func concatInts(num1, num2 *big.Int) *big.Int {
	str := fmt.Sprintf("%s%s", num1.String(), num2.String())
	result := new(big.Int)
	result.SetString(str, 10)
	return result
}

func getDay7Data() [][]*big.Int {
	file, err := os.Open("inputs/input7.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	results := [][]*big.Int{}

	for scanner.Scan() {
		stringRow := strings.Split(scanner.Text(), " ")
		stringRow[0] = strings.Replace(stringRow[0], ":", "", 1)
		intRow := []*big.Int{}

		for _, value := range stringRow {
			bigInt := new(big.Int)
			bigInt.SetString(value, 10)
			intRow = append(intRow, bigInt)
		}
		results = append(results, intRow)
	}
	return results
}
