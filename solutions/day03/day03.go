package day03

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func readInput(filename string) ([][]string, error) {
	inputPath := filepath.Join("inputs", filename)

	file, _ := os.Open(inputPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]string

	cleaner, _ := regexp.Compile(`don\'t\(\).*?do\(\)`)
	pattern, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println(row)

		new_row := cleaner.ReplaceAllString(row, "")
		fmt.Println(new_row)

		matches := pattern.FindAllString(new_row, -1)
		fmt.Println(matches)

		data = append(data, matches)
	}

	return data, nil
}

func SolvePart1(filename string) (int, error) {
	input, _ := readInput(filename)
	fmt.Println(input)

	sum := 0
	digit, _ := regexp.Compile(`\d{1,3}`)

	for _, instr := range input {
		for _, mul := range instr {
			nmbrs := digit.FindAllString(mul, -1)

			i1, _ := strconv.Atoi(nmbrs[0])
			i2, _ := strconv.Atoi(nmbrs[1])

			sum += i1 * i2
		}
	}

	return sum, nil
}

func SolvePart2(filename string) (int, error) {
	input, _ := readInput(filename)
	// fmt.Println(input)

	sum := 0
	digit, _ := regexp.Compile(`\d{1,3}`)

	for _, instr := range input {
		for _, mul := range instr {
			nmbrs := digit.FindAllString(mul, -1)

			i1, _ := strconv.Atoi(nmbrs[0])
			i2, _ := strconv.Atoi(nmbrs[1])

			sum += i1 * i2
		}
	}

	return sum, nil
}
