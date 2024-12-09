package day09

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	row := scanner.Text()

	disk := strings.Split(row, "")

	parsedDisk := []string{}

	fileId := 0
	for idx, digit := range disk {
		file, _ := strconv.Atoi(digit)
		if idx%2 == 0 {
			// digit = file
			for range file {
				fileIdStr := strconv.Itoa(fileId)
				parsedDisk = append(parsedDisk, fileIdStr)
			}
			fileId += 1
		} else {
			// digit = free space
			for range file {
				parsedDisk = append(parsedDisk, ".")
			}
		}
	}

	return parsedDisk
}

func dotsOnly(s []string) bool {
	dot := "."
	for _, val := range s {
		if val != dot {
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
	disk := readInput(filename)

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

func SolvePart2(filename string) int {
	return 0
}
