package day13

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

const EPSILON float64 = 1e-3

type Prize struct {
	x, y float64
}

type Delta struct {
	X, Y float64
}

type Game struct {
	prize  Prize
	deltas map[string]Delta
}

func makeDelta(info []string) Delta {
	X, _ := strconv.Atoi(info[1])
	Y, _ := strconv.Atoi(info[2])

	return Delta{float64(X), float64(Y)}
}

func makeGame(infoA, infoB, prizeInfo []string, pt2 bool) Game {
	deltaA := makeDelta(infoA)
	deltaB := makeDelta(infoB)

	x, _ := strconv.Atoi(prizeInfo[1])
	y, _ := strconv.Atoi(prizeInfo[2])

	incr := 0
	if pt2 {
		incr = 10000000000000
	}

	return Game{
		Prize{float64(x + incr), float64(y + incr)}, map[string]Delta{
			"A": deltaA,
			"B": deltaB,
		},
	}
}

func readInput(filename string, pt2 bool) []Game {
	content, _ := os.ReadFile(filename)

	gameData := strings.Split(string(content), "\n\n")

	patBtn, _ := regexp.Compile(`X\+(\d+), Y\+(\d+)`)
	patPrize, _ := regexp.Compile(`X=(\d+), Y=(\d+)`)

	games := make([]Game, len(gameData))
	for i, game := range gameData {
		input := strings.Split(game, "\n")
		btnA, btnB, prize := input[0], input[1], input[2]

		infoA := patBtn.FindAllStringSubmatch(btnA, -1)[0]
		infoB := patBtn.FindAllStringSubmatch(btnB, -1)[0]
		infoPrize := patPrize.FindAllStringSubmatch(prize, -1)[0]

		games[i] = makeGame(infoA, infoB, infoPrize, pt2)

	}

	return games
}

func isInt(f float64) (bool, int) {
	up := math.Ceil(f)
	down := math.Floor(f)

	if up-f < EPSILON {
		return true, int(up)
	}

	if f-down < EPSILON {
		return true, int(down)
	}

	return false, 0
}

func SolvePart1(filename string) int {
	games := readInput(filename, false)

	c := 0
	for _, game := range games {
		a := mat.NewDense(2, 2, []float64{
			game.deltas["A"].X, game.deltas["B"].X,
			game.deltas["A"].Y, game.deltas["B"].Y,
		},
		)

		var aInv mat.Dense
		err := aInv.Inverse(a)
		if err != nil {
			fmt.Println("A not invertible")
			continue
		}

		y := mat.NewVecDense(2, []float64{game.prize.x, game.prize.y})

		var b mat.VecDense
		b.MulVec(&aInv, y)

		b0, b1 := b.AtVec(0), b.AtVec(1)

		okA, nrA := isInt(b0)
		okB, nrB := isInt(b1)
		if okA && okB {
			c += nrA*3 + nrB
		}

	}

	return c
}

func SolvePart2(filename string) int {
	games := readInput(filename, true)

	c := 0
	for _, game := range games {
		a := mat.NewDense(2, 2, []float64{
			game.deltas["A"].X, game.deltas["B"].X,
			game.deltas["A"].Y, game.deltas["B"].Y,
		},
		)

		var aInv mat.Dense
		err := aInv.Inverse(a)
		if err != nil {
			fmt.Println("A not invertible")
			continue
		}

		y := mat.NewVecDense(2, []float64{game.prize.x, game.prize.y})

		var b mat.VecDense
		b.MulVec(&aInv, y)

		b0, b1 := b.AtVec(0), b.AtVec(1)

		okA, nrA := isInt(b0)
		okB, nrB := isInt(b1)
		if okA && okB {
			c += nrA*3 + nrB
		}

	}

	return c
}
