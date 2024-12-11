package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type TrailHead struct {
	x, y int
}

type Point struct {
	x, y int
}

func (p Point) isLegal(R, C int) bool {
	if p.x < 0 || p.y < 0 || p.x >= R || p.y >= C {
		return false
	}
	return true
}

func (p Point) legalMoves(R, C int) []Point {
	points := []Point{
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
	}

	moves := []Point{}
	for _, point := range points {
		if point.isLegal(R, C) {
			moves = append(moves, point)
		}
	}

	return moves
}

func readInput(filename string) ([][]int, []TrailHead) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	grid := [][]int{}
	trailheads := []TrailHead{}
	for scanner.Scan() {
		heightsStr := strings.Split(scanner.Text(), "")

		j := 0
		heights := []int{}
		for _, s := range heightsStr {
			nr, _ := strconv.Atoi(s)
			heights = append(heights, nr)
			if nr == 0 {
				trailhead := TrailHead{i, j}
				trailheads = append(trailheads, trailhead)

			}
			j += 1
		}
		grid = append(grid, heights)

		i += 1

	}

	return grid, trailheads
}

func pathFinder(p Point, visited map[Point]bool, grid [][]int, pt1 bool) int {
	if pt1 && visited[p] {
		return 0
	}

	visited[p] = true

	if grid[p.x][p.y] == 9 {
		return 1
	}

	R, C := len(grid), len(grid[0])

	pIncr := 0
	for _, pNext := range p.legalMoves(R, C) {
		if grid[p.x][p.y]+1 == grid[pNext.x][pNext.y] {
			pIncr += pathFinder(pNext, visited, grid, pt1)
		}
	}

	return pIncr
}

func SolvePart1(filename string) int {
	grid, trailheads := readInput(filename)

	c := 0
	for _, th := range trailheads {
		c += pathFinder(Point{th.x, th.y}, map[Point]bool{}, grid, true)
	}

	return c
}

func SolvePart2(filename string) int {
	grid, trailheads := readInput(filename)

	c := 0
	for _, th := range trailheads {
		c += pathFinder(Point{th.x, th.y}, map[Point]bool{}, grid, false)
	}

	return c
}
