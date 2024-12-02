package main

import (
	"fmt"
	"log"

	"github.com/Rens660/aoc-2024/solutions/day02"
	//"github.com/Rens660/aoc-2024/solutions/day01"
)

func main() {
	fmt.Println("Advent of Code 2024")

	part1Result, err := day02.SolvePart1("day02.txt")
	if err != nil {
		log.Fatalf("Day02|pt1 error: %v", err)
	}
	fmt.Printf("Day02|pt1: %v\n", part1Result)

	part2Result, err := day02.SolvePart2("day02.txt")
	if err != nil {
		log.Fatalf("Day02|pt2 error: %v", err)
	}
	fmt.Printf("Day02|pt2: %v\n", part2Result)
}
