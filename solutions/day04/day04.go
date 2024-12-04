package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) ([][]string, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()

		data = append(data, strings.Split(row, ""))
	}

	return data, nil
}

func SolvePart1(filename string) int {
	data, _ := readInput(filename)

	for _, row := range data {
		for _, elem := range row {
			fmt.Print(elem)
		}
		fmt.Println()
	}

	return 0
}

func SolvePart2(filename string) int {
	return 0
}
