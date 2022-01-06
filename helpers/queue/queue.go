package queue

type Queue []interface{}

func (q *Queue) Push(i interface{}) {
	*q = append(*q, i)
}

func (q *Queue) Pop() interface{} {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}
