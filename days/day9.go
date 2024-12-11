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
	compressedDiskOptimized := compressedDiskOptimized(diskMap)
	fileSystemCheckSum := calculateChecksum(compressedDisk)
	fileSystemOptimizedCheckSum := calculateChecksum(compressedDiskOptimized)

	log.Printf(
		"Completed disk compression resulted in checksum of %d. An deframented compression has a checksum of %d",
		fileSystemCheckSum,
		fileSystemOptimizedCheckSum,
	)
}

func compressedDiskOptimized(diskMap []int) []int {
	compressedDisk := make([]int, len(diskMap))
	copy(compressedDisk, diskMap)

	getLastGroup := func(fileID int) (start int, size int) {
		lastPos := -1
		for i := len(compressedDisk) - 1; i >= 0; i-- {
			if compressedDisk[i] == fileID {
				lastPos = i
				break
			}
		}
		if lastPos == -1 {
			return -1, 0
		}

		firstPos := lastPos
		for firstPos >= 0 && compressedDisk[firstPos] == fileID {
			firstPos--
		}
		firstPos++

		return firstPos, lastPos - firstPos + 1
	}

	getFirstOpenSpot := func(end int, size int) int {
		for i := 0; i < end; i++ {
			if compressedDisk[i] == -1 {
				spaceSize := 0
				for j := i; j < end && compressedDisk[j] == -1; j++ {
					spaceSize++
					if spaceSize == size {
						return i
					}
				}
			}
		}
		return -1
	}

	for fileID := len(compressedDisk); fileID >= 0; fileID-- {
		start, size := getLastGroup(fileID)
		if start == -1 || size == 0 {
			continue
		}

		newStart := getFirstOpenSpot(start, size)
		if newStart != -1 {
			for i := 0; i < size; i++ {
				compressedDisk[newStart+i] = fileID
			}
			for i := start; i < start+size; i++ {
				compressedDisk[i] = -1
			}
		}
	}

	return compressedDisk
}

func calculateChecksum(compressedDisk []int) int {
	result := 0
	for i, char := range compressedDisk {
		if char == -1 {
			continue
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
