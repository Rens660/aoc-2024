package day03

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func readInput(filename string) (string, error) {
	inputPath := filepath.Join("inputs", filename)

	content, _ := os.ReadFile(inputPath)
	str_content := strings.Replace(string(content), "\n", "", -1)

	return str_content, nil
}

func SolvePart1(filename string) (int, error) {
	input, _ := readInput(filename)

	pattern, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	data := pattern.FindAllString(input, -1)

	sum := 0
	digit, _ := regexp.Compile(`\d{1,3}`)

	for _, mul := range data {
		nmbrs := digit.FindAllString(mul, -1)

		i1, _ := strconv.Atoi(nmbrs[0])
		i2, _ := strconv.Atoi(nmbrs[1])

		sum += i1 * i2
	}

	return sum, nil
}

func SolvePart2(filename string) (int, error) {
	input, _ := readInput(filename)

	cleanupPattern, _ := regexp.Compile(`(don't\(\).*?do\(\))|(don't\(\).*$)`)
	cleaned_input := cleanupPattern.ReplaceAllString(input, "")

	mulPattern, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	data := mulPattern.FindAllString(cleaned_input, -1)

	digitPattern, _ := regexp.Compile(`\d{1,3}`)
	sum := 0

	for _, mul := range data {
		nmbrs := digitPattern.FindAllString(mul, -1)

		i1, _ := strconv.Atoi(nmbrs[0])
		i2, _ := strconv.Atoi(nmbrs[1])

		sum += i1 * i2
	}

	return sum, nil
}
