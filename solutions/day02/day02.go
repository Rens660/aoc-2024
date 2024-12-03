package day02

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")

		var parsed_row []int
		for i := 0; i < len(row); i += 1 {
			value, _ := strconv.Atoi(row[i])
			parsed_row = append(parsed_row, value)
		}

		data = append(data, parsed_row)

	}

	return data, nil
}

func SolvePart1(filename string) int {
	input, _ := readInput(filename)

	nr_safe_reports := 0

	for _, report := range input {
		safe_report := reportIsSafe(report)

		if safe_report {
			nr_safe_reports += 1
		}
	}

	return nr_safe_reports
}

func SolvePart2(filename string) int {
	input, _ := readInput(filename)

	nr_safe_reports := 0

	for _, report := range input {

		safe_report := reportIsSafe(report)
		if safe_report {
			nr_safe_reports += 1
			continue
		}

		j := 0
		for j < len(report) {
			simple_report := removeElementAtIndex(report, j)

			safe_report := reportIsSafe(simple_report)
			if safe_report {
				nr_safe_reports += 1
				break
			}
			j += 1
		}
	}

	return nr_safe_reports
}

func removeElementAtIndex(slice []int, index int) []int {
	result := make([]int, 0, len(slice)-1)

	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)

	return result
}

func reportIsSafe(report []int) bool {
	incr := report[0] < report[1]

	for i := 0; i < len(report)-1; i += 1 {
		diff := int(math.Abs(float64(report[i] - report[i+1])))

		if diff == 0 || diff > 3 {
			return false
		}

		if incr && report[i] >= report[i+1] {
			return false
		}
		if !incr && report[i] <= report[i+1] {
			return false
		}
	}

	return true
}
