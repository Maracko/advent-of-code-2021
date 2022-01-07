package stack

type Stack []interface{}

func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	i := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return i
}
