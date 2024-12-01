package main

import (
	"fmt"
	"log"

	"github.com/Rens660/aoc-2024/solutions/day01"
)

func main() {
	fmt.Println("Advent of Code 2024")

	part1Result, err := day01.SolvePart1("day01.txt")
	if err != nil {
		log.Fatalf("Day01|pt1 error: %v", err)
	}
	fmt.Printf("Day01|pt1: %v\n", part1Result)

	part2Result, err := day01.SolvePart2("day01.txt")
	if err != nil {
		log.Fatalf("Day01|pt2 error: %v", err)
	}
	fmt.Printf("Day01|pt2: %v\n", part2Result)
}
