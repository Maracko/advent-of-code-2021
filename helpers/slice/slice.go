package slice

import (
	"sort"
	"strconv"
	"strings"
)

func SumSliceOfInts(slice []int) int {
	var res int
	for _, i := range slice {
		res += i
	}
	return res
}

func ReadStringToSliceOfInts(s, delimiter string) []int {
	var res []int
	for _, s := range strings.Split(s, delimiter) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}

func ConvertSliceOfStringsToSliceOfInts(slice []string) []int {
	var res []int
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func IntInSlice(slice []int, num int) bool {
	cpy := make([]int, len(slice))
	copy(cpy, slice)
	sort.Ints(cpy)
	i := sort.Search(len(cpy), func(i int) bool { return cpy[i] >= num })
	return i < len(cpy) && cpy[i] == num
}
