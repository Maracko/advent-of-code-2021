package main

import (
	"fmt"
	"strconv"

	"github.com/maracko/advent-of-code-2021/helpers/file"
)

func main() {
	d, _ := file.ReadFileToSliceOfStrings("data.txt")
	matrix := createMatrix(d)

	fmt.Println("Part 1 solution", solvePart1(matrix))
}

func createMatrix(d []string) [][]int {
	var res [][]int
	for _, row := range d {
		line := []int{}
		for _, char := range row {
			num, _ := strconv.Atoi(string(char))
			line = append(line, num)
		}
		res = append(res, line)
	}
	return res
}

func solvePart1(mtx [][]int) int {
	lowPoints := []int{}
	for i, row := range mtx {
		for j, num := range row {
			//Has left which is smaller or equal to current value
			if j > 0 && mtx[i][j-1] <= num {
				continue
			}

			//Has right which is smaller or equal to current value
			if j < len(row)-1 && mtx[i][j+1] <= num {
				continue
			}

			//Has up which ...
			if i > 0 && mtx[i-1][j] <= num {
				continue
			}

			//Has down which ...
			if i < len(mtx)-1 && mtx[i+1][j] <= num {
				continue
			}

			lowPoints = append(lowPoints, num+1)
		}
	}

	res := 0
	for _, point := range lowPoints {
		res += point
	}

	return res
}
