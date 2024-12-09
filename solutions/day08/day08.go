package day08

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

const K = 2

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

func readInput(filename string) ([][]string, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]string{}

	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, strings.Split(row, ""))
	}

	return grid, nil
}

func mapAntennaPositions(grid [][]string) map[string][]Antenna {
	R, C := len(grid), len(grid[0])

	antennaPositions := map[string][]Antenna{}
	for i := range R {
		for j := range C {
			freq := grid[i][j]
			if freq == "." {
				continue
			}

			antenna := Antenna{i, j}
			if antennaPositions[freq] == nil {
				antennaPositions[freq] = []Antenna{antenna}
			} else {
				antennaPositions[freq] = append(antennaPositions[freq], antenna)
			}

		}
	}

	return antennaPositions
}

func (ant Antenna) getAntinodes(ant2 Antenna) (AntiNode, AntiNode) {
	fmt.Println(ant, ant2)

	antinode1 := AntiNode{-1, -1}
	antinode2 := AntiNode{-1, -1}

	dx, dy := ant.deltas(ant2)
	fmt.Println(dx, dy)

	if ant.x == ant2.x {
		if ant.y < ant2.y {
			// add antinode to left of ant and right of ant2
			antinode1 = AntiNode{ant.x, ant.y - dy}
			antinode2 = AntiNode{ant2.x, ant2.y + dy}
		}
		if ant.y > ant2.y {
			// add antinode to right of ant and left of ant2
			antinode1 = AntiNode{ant.x, ant.y + dy}
			antinode2 = AntiNode{ant2.x, ant2.y - dy}
		}
	}

	if ant.x < ant2.x {
		if ant.y == ant2.y {
			// add antinode up of ant and down of ant2
			antinode1 = AntiNode{ant.x - dx, ant.y}
			antinode2 = AntiNode{ant2.x + dx, ant2.y}
		}
		if ant.y < ant2.y {
			// add antinode to upLeft of ant and downRight of ant2
			antinode1 = AntiNode{ant.x - dx, ant.y - dy}
			antinode2 = AntiNode{ant2.x + dx, ant2.y + dy}
		}
		if ant.y > ant2.y {
			// add antinode to upRight of ant and downLeft of ant2
			antinode1 = AntiNode{ant.x - dx, ant.y + dy}
			antinode2 = AntiNode{ant2.x + dx, ant2.y - dy}
		}
	}

	if ant.x > ant2.x {
		if ant.y == ant2.y {
			// add antinode to down of ant and up of ant2
			antinode1 = AntiNode{ant.x + dx, ant.y}
			antinode2 = AntiNode{ant2.x - dx, ant2.y}
		}
		if ant.y < ant2.y {
			// add antinode to downLeft of ant and UpRight of ant2
			antinode1 = AntiNode{ant.x + dx, ant.y - dy}
			antinode2 = AntiNode{ant2.x - dx, ant2.y + dy}
		}
		if ant.y > ant2.y {
			// add antinode to downRight of ant and upLeft of ant2
			antinode1 = AntiNode{ant.x + dx, ant.y + dy}
			antinode2 = AntiNode{ant2.x - dx, ant2.y - dy}
		}
	}

	return antinode1, antinode2
}

func SolvePart1(filename string) int {
	grid, _ := readInput(filename)
	printMatrix(grid)
	R, C := len(grid), len(grid[0])
	antiNodeLedger := map[AntiNode]bool{}

	aps := mapAntennaPositions(grid)
	fmt.Println(aps)

	for _, antennas := range aps {
		N := len(antennas)
		fmt.Println(N)
		if N < K {
			continue
		}

		combis := combin.Combinations(N, K)
		for _, combi := range combis {
			a1, a2 := antennas[combi[0]], antennas[combi[1]]
			node1, node2 := a1.getAntinodes(a2)
			fmt.Println(node1, node2)
			fmt.Println("-----------------\n\n")

			if node1.isLegit(R, C) {
				antiNodeLedger[node1] = true
			}

			if node2.isLegit(R, C) {
				antiNodeLedger[node2] = true
			}
		}
	}

	fmt.Println(antiNodeLedger)

	return len(antiNodeLedger)
}

func (ant Antenna) getAllAntiNodes(ant2 Antenna, ledger map[AntiNode]bool, R, C int) {
	dx, dy := ant.deltas(ant2)
	fmt.Println(dx, dy)

	if ant.x == ant2.x {
		if ant.y < ant2.y {
			// add antinode to left of ant and right of ant2
			ant.storeAntiNodes(0, -dy, R, C, ledger)
			ant2.storeAntiNodes(0, dy, R, C, ledger)
		}

		if ant.y > ant2.y {
			// add antinode to right of ant and left of ant2
			ant.storeAntiNodes(0, dy, R, C, ledger)
			ant2.storeAntiNodes(0, -dy, R, C, ledger)
		}
	}

	if ant.x < ant2.x {
		if ant.y == ant2.y {
			// add antinode up of ant and down of ant2
			ant.storeAntiNodes(-dx, 0, R, C, ledger)
			ant2.storeAntiNodes(dx, 0, R, C, ledger)
		}
		if ant.y < ant2.y {
			// add antinode to upLeft of ant and downRight of ant2
			ant.storeAntiNodes(-dx, -dy, R, C, ledger)
			ant2.storeAntiNodes(dx, dy, R, C, ledger)
		}
		if ant.y > ant2.y {
			// add antinode to upRight of ant and downLeft of ant2
			ant.storeAntiNodes(-dx, dy, R, C, ledger)
			ant2.storeAntiNodes(dx, -dy, R, C, ledger)
		}
	}

	if ant.x > ant2.x {
		if ant.y == ant2.y {
			// add antinode to down of ant and up of ant2
			ant.storeAntiNodes(dx, 0, R, C, ledger)
			ant2.storeAntiNodes(-dx, 0, R, C, ledger)
		}
		if ant.y < ant2.y {
			// add antinode to downLeft of ant and UpRight of ant2
			ant.storeAntiNodes(dx, -dy, R, C, ledger)
			ant2.storeAntiNodes(-dx, dy, R, C, ledger)
		}
		if ant.y > ant2.y {
			// add antinode to downRight of ant and upLeft of ant2
			ant.storeAntiNodes(dx, dy, R, C, ledger)
			ant2.storeAntiNodes(-dx, -dy, R, C, ledger)
		}
	}
}

func SolvePart2(filename string) int {
	grid, _ := readInput(filename)
	R, C := len(grid), len(grid[0])

	aps := mapAntennaPositions(grid)

	ledger := map[AntiNode]bool{}
	for _, antennas := range aps {
		N := len(antennas)
		if N < K {
			continue
		}

		combis := combin.Combinations(N, K)
		for _, combi := range combis {
			a1, a2 := antennas[combi[0]], antennas[combi[1]]

			a1Node := AntiNode{a1.x, a1.y}
			ledger[a1Node] = true

			a2Node := AntiNode{a2.x, a2.y}
			ledger[a2Node] = true

			a1.getAllAntiNodes(a2, ledger, R, C)
		}

	}

	return len(ledger)
}
