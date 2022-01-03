package math

import "math"

func IntAbs(num int) int {
	return int(math.Abs((float64(num))))
}

func SumNaturalsInRange(n int) int {
	return (n * (n + 1)) / 2
}

func GetMeanAverage(nums []int) int {
	sum := 0.0
	for _, num := range nums {
		sum += float64(num)
	}
	return int(math.Floor(sum / float64(len(nums))))
}
