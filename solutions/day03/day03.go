package day03

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(filename string) (string, error) {
	content, _ := os.ReadFile(filename)
	str_content := strings.Replace(string(content), "\n", "", -1)

	return str_content, nil
}

func SolvePart1(filename string) int {
	input, _ := readInput(filename)

	pattern, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	data := pattern.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, mul := range data {

		i1, _ := strconv.Atoi(mul[1])
		i2, _ := strconv.Atoi(mul[2])

		sum += i1 * i2
	}

	return sum
}

func SolvePart2(filename string) int {
	input, _ := readInput(filename)

	cleanupPattern, _ := regexp.Compile(`(don't\(\).*?do\(\))|(don't\(\).*$)`)
	cleaned_input := cleanupPattern.ReplaceAllString(input, "")

	mulPattern, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	data := mulPattern.FindAllStringSubmatch(cleaned_input, -1)

	sum := 0
	for _, mul := range data {
		i1, _ := strconv.Atoi(mul[1])
		i2, _ := strconv.Atoi(mul[2])
		sum += i1 * i2
	}

	return sum
}
