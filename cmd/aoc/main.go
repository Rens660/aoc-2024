package main

import (
	"fmt"

	solver "github.com/Rens660/aoc-2024/solutions/day09"
)

type AOC struct {
	day  int
	part int
	test bool
}

func (aoc AOC) Input() string {
	var day string
	if aoc.day < 10 {
		day = fmt.Sprintf("day0%d", aoc.day)
	} else {
		day = fmt.Sprintf("day%d", aoc.day)
	}

	if aoc.test {
		return fmt.Sprintf("inputs/test/%v_test.txt", day)
	}

	return fmt.Sprintf("inputs/%v.txt", day)
}

func (aoc AOC) solve() (int, int) {
	inputFilePath := aoc.Input()

	switch aoc.part {
	case 1:
		pt1 := solver.SolvePart1(inputFilePath)
		return pt1, -404
	case 2:
		pt2 := solver.SolvePart2(inputFilePath)
		return -404, pt2
	default:
		pt1 := solver.SolvePart1(inputFilePath)
		pt2 := solver.SolvePart2(inputFilePath)
		return pt1, pt2
	}
}

func main() {
	fmt.Println(`
 ╔══════════════════════════════════════╗
 ║          ADVENT OF CODE 2024         ║
 ║    Unwrap a Coding Challenge Today!  ║
 ╚══════════════════════════════════════╝
      🎁 Code · Solve · Celebrate 🎁

        LETS SOLVE SOME PUZZLES!
  --------------------------------------
  `)

	aoc := AOC{day: 9, part: 1, test: false}

	pt1, pt2 := aoc.solve()

	fmt.Printf("Day %d | pt1: %v\n", aoc.day, pt1)
	fmt.Printf("Day %d | pt2: %v\n", aoc.day, pt2)
}
