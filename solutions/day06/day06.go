package day06

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func readInput(filename string) ([][]string, error) {
	file, _ := os.Open(filename)
	defer file.Close()

	var room [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()

		room = append(room, strings.Split(row, ""))
	}

	return room, nil
}

func initStartCoordinate(room [][]string, nrRows int, nrCols int) (Coordinate, error) {
	for i := range nrRows {
		for j := range nrCols {
			if room[i][j] == "^" {
				return Coordinate{i, j}, nil
			}
		}
	}

	return Coordinate{}, errors.New("Where is the guard?!")
}

func makePath(room [][]string, guard *Guard) []Coordinate {
	var path []Coordinate
	set := make(map[Coordinate]bool)

	for {
		if !set[guard.Coordinate] {
			set[guard.Coordinate] = true
			path = append(path, guard.Coordinate)
		}

		err := guard.move(room)
		if err != nil {
			break
		}
		if guard.InLoop {
			return []Coordinate{}
		}

	}

	return path
}

func SolvePart1(filename string) int {
	room, _ := readInput(filename)

	nrRows, nrCols := len(room), len(room[0])
	coordinate, _ := initStartCoordinate(room, nrRows, nrCols)
	guard := Guard{
		Heading:    headingMapper[1],
		Coordinate: coordinate,
		Blockades:  map[Coordinate]string{},
	}

	path := makePath(room, &guard)
	return len(path)
}

func SolvePart2(filename string) int {
	room, _ := readInput(filename)
	start := time.Now()

	nrRows, nrCols := len(room), len(room[0])
	coordinate, _ := initStartCoordinate(room, nrRows, nrCols)
	guard := Guard{
		Heading:    headingMapper[1],
		Coordinate: coordinate,
		Blockades:  map[Coordinate]string{},
		InLoop:     false,
	}

	path := makePath(room, &guard)

	nrLoopIntroducingObstacles := 0

	for _, node := range path[1:] {
		guard = Guard{
			Heading:    headingMapper[1],
			Coordinate: coordinate,
			Blockades:  map[Coordinate]string{},
			InLoop:     false,
		}

		room[node.x][node.y] = "#"

		makePath(room, &guard)
		if guard.InLoop {
			nrLoopIntroducingObstacles += 1
		}

		room[node.x][node.y] = "."
	}

	t := time.Now()
	elapsed := t.Sub(start).Seconds()
	fmt.Printf("Time elapsed: %f\n", elapsed)

	return nrLoopIntroducingObstacles
}
