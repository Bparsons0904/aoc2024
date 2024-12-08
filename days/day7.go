package days

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func Day7() {
	data := getDay7Data()
	calibrationTotals := big.NewInt(0)
	calibrationsToRetotal := [][]*big.Int{}

	for _, calibrations := range data {
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
		found := false
		for _, result := range potentials {
			if result.Cmp(expectedTotal) == 0 {
				calibrationTotals.Add(calibrationTotals, result)
				found = true
				break
			}
		}
		if !found {
			calibrationsToRetotal = append(calibrationsToRetotal, calibrations)
		}
	}

	expandedCalibrationsTotal := big.NewInt(0)
	for _, calibrations := range calibrationsToRetotal {
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
				expandedCalibrationsTotal.Add(expandedCalibrationsTotal, result)
				break
			}
		}
	}

	totalSum := new(big.Int).Add(calibrationTotals, expandedCalibrationsTotal)
	log.Printf(
		"Total of calibrations: %s with an expanded Calibration Total of: %s",
		calibrationTotals.String(),
		totalSum.String(),
	)
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
