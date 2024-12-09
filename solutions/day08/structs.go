package day08

import (
	"math"
)

type Antenna struct {
	x int
	y int
}

func (ant Antenna) deltas(ant2 Antenna) (int, int) {
	dx := math.Abs(float64(ant.x - ant2.x))
	dy := math.Abs(float64(ant.y - ant2.y))

	return int(dx), int(dy)
}

type AntiNode struct {
	x int
	y int
}

func (ant Antenna) next(dx, dy int) AntiNode {
	antinode := AntiNode{ant.x + dx, ant.y + dy}
	return antinode
}

func (an AntiNode) isLegit(R, C int) bool {
	if an.x < 0 || an.y < 0 {
		return false
	}

	if an.x >= R || an.y >= C {
		return false
	}

	return true
}

func (ant Antenna) storeAntiNodes(diffX, diffY, R, C int, ledger map[AntiNode]bool) {
	k := 1
	for {
		antinode := ant.next(k*diffX, k*diffY)
		if antinode.isLegit(R, C) {
			ledger[antinode] = true
			k += 1
		} else {
			break
		}
	}
}
