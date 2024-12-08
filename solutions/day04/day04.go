package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes)
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

type XY struct {
	x int
	y int
}

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

	R, C := len(data), len(data[0])

	joinedRows := make([]string, R)
	for k, row := range data {
		t := strings.Join(row, "")
		joinedRows[k] = t

	}

	joinedCols := transpose(data, R, C)
	joinedDiags1 := extractDiagonalsNorthEast(data, R, C)
	joinedDiags2 := extractDiagonalsNorthWest(data, R, C)

	joined := [][]string{
		joinedRows,
		joinedCols,
		joinedDiags1,
		joinedDiags2,
	}

	c := 0
	pat, _ := regexp.Compile(`XMAS`)

	for _, data := range joined {
		for _, text := range data {

			matches := pat.FindAllString(text, -1)
			c += len(matches)

			matches = pat.FindAllString(Reverse(text), -1)
			c += len(matches)

		}
	}
	return c
}

func SolvePart2(filename string) int {
	data, _ := readInput(filename)

	R, C := len(data), len(data[0])

	c := 0

	for i := 1; i < R-1; i++ {
		for j := 1; j < C-1; j++ {
			if data[i][j] != "A" {
				continue
			}

			leftUp := XY{i - 1, j - 1}
			rightUp := XY{i - 1, j + 1}
			leftDown := XY{i + 1, j - 1}
			rightDown := XY{i + 1, j + 1}

			stringOne := strings.Join(
				[]string{
					data[leftUp.x][leftUp.y],
					"A",
					data[rightDown.x][rightDown.y],
				},
				"",
			)

			stringTwo := strings.Join(
				[]string{
					data[leftDown.x][leftDown.y],
					"A",
					data[rightUp.x][rightUp.y],
				},
				"",
			)

			if (stringOne == "MAS" ||
				stringOne == "SAM") &&
				(stringTwo == "MAS" ||
					stringTwo == "SAM") {
				c += 1
			}

			// fmt.Println(stringOne)
			// fmt.Println(stringTwo)

		}
	}

	return c
}
