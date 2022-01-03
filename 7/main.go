package main

import (
	"fmt"
	"sort"

	"github.com/maracko/advent-of-code-2021/helpers/file"
	hMath "github.com/maracko/advent-of-code-2021/helpers/math"
	"github.com/maracko/advent-of-code-2021/helpers/slice"
)

func main() {

	d, _ := file.ReadFileToSliceOfStrings("data.txt")
	data := slice.ReadStringToSliceOfInts(d[0], ",")
	sort.Ints(data)

	fmt.Println("Part 1 solution:", solvePart1(data))
	fmt.Println("Part 2 solution:", solvePart2(data))

}

func solvePart1(nums []int) int {
	median := nums[len(nums)/2]
	fuel := 0

	for _, num := range nums {
		fuel += hMath.IntAbs(median - num)
	}
	return fuel
}

func solvePart2(nums []int) int {

	res := 0
	mean := hMath.GetMeanAverage(nums)

	for _, num := range nums {
		distance := hMath.IntAbs(num - mean)
		fuel := hMath.SumNaturalsInRange(distance)
		res += fuel
	}

	return res
}
