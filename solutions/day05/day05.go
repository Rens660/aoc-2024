package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

const K = 2

func readInput(filename string) (map[int]map[int]bool, [][]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	rules := map[int]map[int]bool{}
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	ruleMode := true

	for scanner.Scan() {

		row := scanner.Text()
		if row == "" {
			ruleMode = false
			continue
		}

		if ruleMode {
			rule := strings.Split(row, "|")

			key, _ := strconv.Atoi(rule[0])
			value, _ := strconv.Atoi(rule[1])

			if rules[key] == nil {
				rules[key] = map[int]bool{value: true}
			} else {
				rules[key][value] = true
			}

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

	c := 0
	for _, update := range updates {
		legalUpdate := true

		N := len(update)

		combis := combin.Combinations(N, K)
		for _, combi := range combis {
			a, b := update[combi[1]], update[combi[0]]
			if rules[a][b] {
				legalUpdate = false
				break
			}
		}

		if legalUpdate {
			middleIdx := N / 2
			c += update[middleIdx]
		}

	}

	return c
}

func SolvePart2(filename string) int {
	rules, updates := readInput(filename)

	c := 0
	for _, update := range updates {
		illegalUpdate := false

		N := len(update)

		combis := combin.Combinations(N, K)

		k := 0
		for k < len(combis) {
			for _, combi := range combis {
				a, b := update[combi[1]], update[combi[0]]
				if rules[a][b] {
					if !illegalUpdate {
						illegalUpdate = true
					}

					update[combi[0]] = a
					update[combi[1]] = b

					k = 0

				} else {
					k += 1
				}
			}
		}

		if illegalUpdate {
			middleIdx := N / 2
			c += update[middleIdx]
		}

	}

	return c
}
