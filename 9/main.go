package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/maracko/advent-of-code-2021/helpers/file"
	"github.com/maracko/advent-of-code-2021/helpers/queue"
	"github.com/maracko/advent-of-code-2021/helpers/slice"
)

type point struct {
	row, col int
}

func main() {
	d, _ := file.ReadFileToSliceOfStrings("data.txt")
	matrix := createMatrix(d)

	p1Res, lowestPoints := solvePart1(matrix)
	fmt.Println("Part 1 solution", p1Res)
	fmt.Println("Part 2 soltion", solvePart2(matrix, lowestPoints))
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

func solvePart1(mtx [][]int) (int, []point) {
	riskLevels := []int{}
	lowPoints := []point{}
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

			// P1 res
			riskLevels = append(riskLevels, num+1)
			// Slice of lowPoints for P2
			lowPoints = append(lowPoints, point{i, j})
		}
	}

	return slice.SumSliceOfInts(riskLevels), lowPoints
}

func (p point) getBasinSize(mtx [][]int) int {
	size := 0
	visited := createVisitedSlice(mtx)
	queue := queue.Queue{p}

	// Perform a breadth first traversal through the matrix
	for len(queue) > 0 {
		current := queue.Pop().(point)

		if outOfBounds := current.row < 0 || current.col < 0 || current.row >= len(mtx) || current.col >= len(mtx[current.row]); outOfBounds {
			continue
		}
		if alreadyVisited := visited[current.row][current.col]; alreadyVisited {
			continue
		}
		if tooHigh := mtx[current.row][current.col] == 9; tooHigh {
			continue
		}

		visited[current.row][current.col] = true
		size++

		queue.Push(point{current.row, current.col - 1}) // Go left
		queue.Push(point{current.row, current.col + 1}) // Go right
		queue.Push(point{current.row - 1, current.col}) // Go up
		queue.Push(point{current.row + 1, current.col}) // Go down

	}

	return size
}

func createVisitedSlice(mtx [][]int) [][]bool {
	visited := make([][]bool, len(mtx))
	for i := range visited {
		visited[i] = make([]bool, len(mtx[i]))
	}
	return visited
}

func solvePart2(mtx [][]int, pts []point) int {
	largestBasinSizes := []int{}

	for _, point := range pts {
		size := point.getBasinSize(mtx)
		largestBasinSizes = append(largestBasinSizes, size)
	}

	sort.Ints(largestBasinSizes)
	highest := largestBasinSizes[len(largestBasinSizes)-3:]
	res := 1
	for _, num := range highest {
		res *= num
	}
	return res
}
