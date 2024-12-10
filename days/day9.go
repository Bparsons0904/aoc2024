package days

import (
	"bufio"
	"log"
	"os"
)

func Day9() {
	data := getDay9Data()

	diskMap := getDiskMap(data)
	compressedDisk := compressDisk(diskMap)
	fileSystemCheckSum := calculateChecksum(compressedDisk)

	log.Println(fileSystemCheckSum)
}

func calculateChecksum(compressedDisk []int) int {
	result := 0
	for i, char := range compressedDisk {
		if char == -1 {
			break
		}
		result += i * int(char)
	}
	return result
}

func compressDisk(diskMap []int) []int {
	compressedDisk := make([]int, len(diskMap))
	copy(compressedDisk, diskMap)
	for i := 0; i < len(compressedDisk); i++ {
		if compressedDisk[i] == -1 {
			for j := len(compressedDisk) - 1; j > i; j-- {
				if compressedDisk[j] != -1 {
					compressedDisk[i] = compressedDisk[j]
					compressedDisk[j] = -1
					break
				}
			}
		}
	}

	return compressedDisk
}

func getDiskMap(data string) []int {
	diskMap := []int{}
	id := 0
	for i, idChar := range data {
		if i%2 == 0 {
			for j := 0; j < int(idChar-'0'); j++ {
				diskMap = append(diskMap, id)
			}
			id++
		} else {
			for j := 0; j < int(idChar-'0'); j++ {
				diskMap = append(diskMap, -1)
			}
		}
	}

	return diskMap
}

func getDay9Data() string {
	file, err := os.Open("inputs/input9.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
