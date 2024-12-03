package main

import (
	"fmt"
	"log"

	"github.com/Rens660/aoc-2024/solutions/day03"
	//"github.com/Rens660/aoc-2024/solutions/day01"
	//"github.com/Rens660/aoc-2024/solutions/day01"
)

func main() {
	fmt.Println("Advent of Code 2024")

	part1Result, err := day03.SolvePart1("day03.txt")
	if err != nil {
		log.Fatalf("Day03|pt1 error: %v", err)
	}

	fmt.Printf("Day03|pt1: %v\n", part1Result)

	part2Result, err := day03.SolvePart2("day03.txt")
	if err != nil {
		log.Fatalf("Day03|pt2 error: %v", err)
	}
	fmt.Printf("Day03|pt2: %v\n", part2Result)
}
