package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func nrDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

type Stone struct {
	value, length int
}

func NewStone(value int) Stone {
	stone := Stone{value, 0}
	stone.length += nrDigits(value)

	return stone
}

func (s Stone) Morph() []Stone {
	if s.value == 0 {
		return []Stone{NewStone(1)}
	}
	if s.length%2 == 0 {
		lengthAsStr := strconv.Itoa(s.value)

		left, _ := strconv.Atoi(lengthAsStr[:s.length/2])
		right, _ := strconv.Atoi(lengthAsStr[s.length/2:])

		return []Stone{NewStone(left), NewStone(right)}
	}

	return []Stone{NewStone(2024 * s.value)}
}

func readInput(filename string) []Stone {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stones := []Stone{}
	for _, s := range strings.Split(scanner.Text(), " ") {
		val, _ := strconv.Atoi(s)
		stone := NewStone(val)
		stones = append(stones, stone)
	}

	return stones
}

func blinkMagic(stones []Stone) []Stone {
	stonesAfterBlink := []Stone{}

	for _, stone := range stones {
		newStones := stone.Morph()

		stonesAfterBlink = append(stonesAfterBlink, newStones[0])
		if len(newStones) == 2 {
			stonesAfterBlink = append(stonesAfterBlink, newStones[1])
		}

	}

	return stonesAfterBlink
}

func SolvePart1(filename string) int {
	stones := readInput(filename)

	blinkCount := 1

	for blinkCount <= 25 {
		stones = blinkMagic(stones)
		blinkCount++
	}

	return len(stones)
}

func SolvePart2(filename string) int {
	stones := readInput(filename)

	stoneMap := map[Stone]int{}
	for _, stone := range stones {
		stoneMap[stone] += 1
	}

	var blinkedStones map[Stone]int

	blinks := 1
	for blinks <= 75 {
		blinkedStones = map[Stone]int{}

		for stone, count := range stoneMap {
			delete(stoneMap, stone)

			newStones := stone.Morph()
			blinkedStones[newStones[0]] += count
			if len(newStones) == 2 {
				blinkedStones[newStones[1]] += count
			}

		}

		for newBlinkedStoneValue, count := range blinkedStones {
			stoneMap[newBlinkedStoneValue] += count
		}

		blinks++

	}

	c := 0
	for _, stoneCount := range blinkedStones {
		c += stoneCount
	}

	return c
}
