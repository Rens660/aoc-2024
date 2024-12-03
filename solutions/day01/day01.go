package day01

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Rens660/aoc-2024/pkg/utils"
)

func readInput(filename string, pt1 bool) ([][]int, int, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	size := 0

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "   ")

		i, _ := strconv.Atoi(row[0])
		left = append(left, i)

		j, _ := strconv.Atoi(row[1])
		right = append(right, j)

		size += 1

	}

	if pt1 {
		utils.SortAsc(&left)
		utils.SortAsc(&right)
	}

	return [][]int{left, right}, size, nil
}

func SolvePart1(filename string) int {
	input, size, _ := readInput(filename, true)

	i := 0
	c := 0

	for i < size {
		abs_diff := math.Abs(float64(input[0][i] - input[1][i]))

		c += int(abs_diff)
		i += 1
	}

	return c
}

func SolvePart2(filename string) int {
	input, _, _ := readInput(filename, false)

	left, right := input[0], input[1]
	rightMap := utils.MakeMap(&right)

	c := 0

	for _, leftVal := range left {
		c += leftVal * rightMap[leftVal]
	}

	return c
}
