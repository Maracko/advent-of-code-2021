package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line [5]int
type board [5]line

var numbers []int

func main() {
	f, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	n := strings.Split(lines[0], ",")
	numbers = make([]int, len(n))
	for i, v := range n {
		numbers[i], _ = strconv.Atoi(v)
	}

	var boards []board
	var b board
	skipped := 0
	for i, line := range lines[2:] {
		if line == "" {
			skipped++
			continue
		}
		var rowIdx int
		if i < 5 {
			rowIdx = i
		} else {
			rowIdx = (i - skipped) % 5
		}

		if rowIdx == 0 {
			b = board{}
		}

		col := 0
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			b[rowIdx][col], _ = strconv.Atoi(num)
			col += 1
		}

		if rowIdx == 4 {
			boards = append(boards, b)
		}

	}

	lastWinningBoard := board{}
	lastDrawnNumIdx := 0
	for i := range numbers {
		nDrawn := i + 1
		for {
			j := 0
			for boardIdx, board := range boards {
				j++
				rows, columns := getAllRowsAndColumns(board)
				winnerBoard := false

				if won := lineWon(nDrawn, rows); won {
					winnerBoard = true
				}
				if won := lineWon(nDrawn, columns); won {
					winnerBoard = true
				}

				if winnerBoard {
					lastWinningBoard = board
					lastDrawnNumIdx = i
					boards = append(boards[:boardIdx], boards[boardIdx+1:]...)
					j = 0
					break
				}
			}
			if j == len(boards) {
				break
			}

		}
	}
	fmt.Printf("Part 2 solution: %d\n", sumAllRemainingNumbers(lastDrawnNumIdx+1, lastWinningBoard)*numbers[lastDrawnNumIdx])
}

// Sum all reamining numbers on board after removing every number that has been drawn
func sumAllRemainingNumbers(nDrawn int, b board) int {
	var sum int
	for _, row := range b {
		for _, num := range row {
			if !numInArray(num, numbers[:nDrawn]) {
				sum += num
			}
		}
	}
	return sum
}

func lineWon(nDrawn int, lns [5]line) bool {
	for _, ln := range lns {
		if isWinningLine(nDrawn, ln) {
			return true
		}
	}
	return false
}

func getAllRowsAndColumns(b board) ([5]line, [5]line) {
	var rows [5]line
	var columns [5]line

	for rowIdx, row := range b {
		rows[rowIdx] = row
		for colIdx, num := range row {
			columns[colIdx][rowIdx] = num
		}

	}

	return rows, columns
}

func isWinningLine(nDrawn int, ln line) bool {
	if nDrawn < 5 {
		return false
	}

	nMatched := 0
	for _, num := range ln {
		if numInArray(num, numbers[:nDrawn]) {
			nMatched++
		}
	}

	return nMatched == 5
}

func numInArray(num int, arr []int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}
