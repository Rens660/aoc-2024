package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) map[int][][]int {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := map[int][][]int{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ": ")

		testVal, _ := strconv.Atoi(row[0])

		nmbrs := []int{}
		for _, s := range strings.Split(row[1], " ") {
			nr, _ := strconv.Atoi(s)
			nmbrs = append(nmbrs, nr)
		}

		if m[testVal] == nil {
			m[testVal] = [][]int{}
		}
		m[testVal] = append(m[testVal], nmbrs)

	}

	return m
}

func possibleTestValue(testValue, acc, k int, nmbrs []int, addConcat bool) bool {
	if k == len(nmbrs)-1 {
		return acc == testValue
	}

	add := possibleTestValue(testValue, acc+nmbrs[k+1], k+1, nmbrs, addConcat)
	product := possibleTestValue(testValue, acc*nmbrs[k+1], k+1, nmbrs, addConcat)
	cnct := false
	if addConcat {
		joined := fmt.Sprintf("%d%d", acc, nmbrs[k+1])
		joinedInt, _ := strconv.Atoi(joined)

		cnct = possibleTestValue(testValue, joinedInt, k+1, nmbrs, addConcat)
	}

	return add || product || cnct
}

func SolvePart1(filename string) int {
	input := readInput(filename)

	var c int = 0
	for tv, nmbrsL := range input {
		for _, nmbrs := range nmbrsL {
			if possibleTestValue(tv, nmbrs[0], 0, nmbrs, false) {
				c += tv
			}
		}
	}

	return c
}

func SolvePart2(filename string) int {
	input := readInput(filename)

	var c int = 0
	for tv, nmbrsL := range input {
		for _, nmbrs := range nmbrsL {
			if possibleTestValue(tv, nmbrs[0], 0, nmbrs, true) {
				c += tv
			}
		}
	}

	return c
}
