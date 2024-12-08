package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (map[int][]int, [][]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	rules := map[int][]int{}
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	ruleMode := true

	for scanner.Scan() {

		row := scanner.Text()
		if row == "" {
			ruleMode = false
			continue
		}
		fmt.Println(row)

		if ruleMode {
			rule := strings.Split(row, "|")

			key, _ := strconv.Atoi(rule[0])
			value, _ := strconv.Atoi(rule[1])

			rules[key] = append(rules[key], value)

		} else {
			rawPage := strings.Split(row, ",")

			page := make([]int, len(rawPage))

			for i, ch := range rawPage {
				pageNr, _ := strconv.Atoi(ch)
				page[i] = pageNr
			}

			updates = append(updates, page)

		}

	}

	return rules, updates
}

func SolvePart1(filename string) int {
	rules, updates := readInput(filename)

	fmt.Println(rules)
	fmt.Println(updates)

	return 0
}

func SolvePart2(filename string) int {
	// data, _ := readInput(filename)
	return 0
}
