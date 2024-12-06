package day06

import (
	"errors"
)

type Coordinate struct {
	x int
	y int
}

type Heading struct {
	id        int
	direction string
	dx        int
	dy        int
}

var headingMapper = map[int]Heading{
	1: {1, "NORTH", -1, 0},
	2: {2, "EAST", 0, 1},
	3: {3, "SOUTH", 1, 0},
	4: {4, "WEST", 0, -1},
}

type Guard struct {
	Coordinate Coordinate
	Heading    Heading
	Blockades  map[Coordinate]string
	InLoop     bool
}

func (g *Guard) rotate90CW() {
	next := g.Heading.id + 1
	if next > 4 {
		next = 1
	}

	g.Heading = headingMapper[next]
}

func (g *Guard) move(room [][]string) error {
	new_x := g.Coordinate.x + g.Heading.dx
	new_y := g.Coordinate.y + g.Heading.dy

	newCoordinate := Coordinate{new_x, new_y}

	if new_x < 0 || new_y < 0 || new_x >= len(room) || new_y >= len(room[0]) {
		return errors.New("Guard left the room")
	}

	if room[new_x][new_y] == "#" {
		if g.Blockades[newCoordinate] == g.Heading.direction {
			g.InLoop = true
			return nil
		}

		g.Blockades[newCoordinate] = g.Heading.direction
		g.rotate90CW()

	} else {
		g.Coordinate = newCoordinate
	}

	return nil
}
