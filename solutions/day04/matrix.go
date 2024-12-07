package day04

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func transpose(matrix [][]string, R int, C int) []string {
	matrixT := make([][]string, C)
	for i := range matrixT {
		matrixT[i] = make([]string, R)
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			matrixT[j][i] = matrix[i][j]
		}
	}

	joined := make([]string, C)
	for k, row := range matrixT {
		t := strings.Join(row, "")
		joined[k] = t
	}

	return joined
}

func extractDiagonalsNorthEast(matrix [][]string, R, C int) []string {
	diagonals := make([]string, R+C-1)

	for k := 0; k < R+C-1; k++ {
		diagonal := []string{}

		startRow := min(k, R-1)
		startCol := max(0, k-R+1)

		for startRow >= 0 && startCol < C {
			diagonal = append(diagonal, matrix[startRow][startCol])
			startRow--
			startCol++
		}

		t := strings.Join(diagonal, "")
		diagonals[k] = t

	}

	return diagonals
}

func extractDiagonalsNorthWest(matrix [][]string, R, C int) []string {
	fmt.Printf("R: %d\nC: %d\n", R, C)

	diagonals := make([]string, R+C-1)

	for k := 0; k < R+C-1; k++ {
		diagonal := []string{}

		startRow := max(0, k-C+1)
		startCol := max(C-1-k, 0)

		for startRow < R && startCol < C {
			diagonal = append(diagonal, matrix[startRow][startCol])
			startRow++
			startCol++
		}

		t := strings.Join(diagonal, "")
		diagonals[k] = t
	}

	return diagonals
}
