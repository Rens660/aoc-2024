package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]string, map[int][]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	row := scanner.Text()

	disk := strings.Split(row, "")

	parsedDisk := []string{}
	fileIdsMap := map[int][]int{}

	fileId := 0
	for idx, digit := range disk {
		file, _ := strconv.Atoi(digit)
		if idx%2 == 0 {
			// digit = file
			fileIdIndices := []int{}
			for range file {
				fileIdStr := strconv.Itoa(fileId)
				parsedDisk = append(parsedDisk, fileIdStr)
				fileIdIndices = append(fileIdIndices, len(parsedDisk)-1)
			}

			fileIdsMap[fileId] = fileIdIndices
			fileId += 1

		} else {
			// digit = free space
			for range file {
				parsedDisk = append(parsedDisk, ".")
			}
		}
	}

	return parsedDisk, fileIdsMap
}

func dotsOnly(s []string) bool {
	for _, val := range s {
		if val != "." {
			return false
		}
	}

	return true
}

func findNextDot(idx int, s []string) int {
	for k, elem := range s[idx:] {
		if elem == "." {
			return k + idx
		}
	}

	return 0
}

func SolvePart1(filename string) int {
	disk, _ := readInput(filename)

	i := 0
	for k := len(disk) - 1; k >= 0; k-- {
		if disk[k] == "." {
			continue
		}

		i = findNextDot(i+1, disk)

		afterDot := disk[i:]
		if dotsOnly(afterDot) {
			break
		}

		disk[i] = disk[k]
		disk[k] = "."

	}

	c := 0
	for j, file := range disk {
		if file == "." {
			break
		}

		nr, _ := strconv.Atoi(file)
		c += j * nr

	}

	return c
}

func findNextDots(nrDotsNeeded int, disk []string) []int {
	closestDotIndices := []int{}

	for i := 0; i <= len(disk)-nrDotsNeeded; i++ {
		x := disk[i : i+nrDotsNeeded]

		if dotsOnly(x) {
			for k := range nrDotsNeeded {
				closestDotIndices = append(closestDotIndices, i+k)
			}
			break
		}
	}

	return closestDotIndices
}

func SolvePart2(filename string) int {
	disk, fileIdMap := readInput(filename)

	maxFileId := len(fileIdMap)

	for fileId := maxFileId; fileId >= 0; fileId-- {
		indices := fileIdMap[fileId]
		spaceNeeded := len(indices)

		nextDots := findNextDots(spaceNeeded, disk)
		if len(nextDots) == 0 {
			continue
		}

		if nextDots[0] > indices[0] {
			continue
		}

		fileIdStr := strconv.Itoa(fileId)

		for _, p := range nextDots {
			disk[p] = fileIdStr
		}

		for _, q := range indices {
			disk[q] = "."
		}
	}

	c := 0
	for j, file := range disk {
		if file == "." {
			continue
		}

		nr, _ := strconv.Atoi(file)
		c += j * nr

	}

	return c
}
