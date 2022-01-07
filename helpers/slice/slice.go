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

func LowerNumberExistsInSlice(s []int, n int) (exists bool, idx int) {
	for i := 0; i < len(s); i++ {
		num := s[i]
		if num < n {
			exists = true
			idx = i
			return

		}
	}
	return
}

func RuneInSlice(s []rune, i rune) (exists bool, idx int) {
	for idx, val := range s {
		if val == i {
			return true, idx
		}
	}
	return false, -1
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func SortAllStringsInSlice(s []string) []string {
	var res []string
	for _, str := range s {
		runes := []rune(str)
		sort.Sort(sortRuneString(runes))
		res = append(res, string(runes))
	}
	return res
}
