package main

import (
	"fmt"
	"strconv"

	"github.com/maracko/advent-of-code-2021/helpers"
)

func main() {

	vals, _ := helpers.ReadFileToSliceOfStrings("data.txt")
	intVals := make([]int, 0, len(vals))
	for _, str := range vals {
		val, _ := strconv.Atoi(str)
		intVals = append(intVals, val)
	}

	fmt.Println("Part 1 solution:", solvePart1(intVals))
	fmt.Println("Part 2 solution:", solvePart2(intVals))

}

func solvePart1(nums []int) int {
	lastVal := 0
	i := 0
	timesHigher := 0
	for _, val := range nums {
		i++
		if i == 1 {
			lastVal = val
			continue
		}
		if val > lastVal {
			timesHigher++
		}
		lastVal = val
	}

	return timesHigher
}

func solvePart2(nums []int) int {
	timesIncreased := 0
	for i := range nums {
		if i+3 >= len(nums) {
			break
		}

		if (nums[i] + nums[i+1] + nums[i+2]) < (nums[i+1] + nums[i+2] + nums[i+3]) {
			timesIncreased++
		}
	}
	return timesIncreased
}
