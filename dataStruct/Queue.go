package dataStruct

type Queue struct {
	container []int
}

func NewQueue() Queue {
	return Queue{
		make([]int, 0),
	}
}
func (q *Queue) Size() int {
	return len(q.container)
}
func (q *Queue) Enqueue(x int) {
	q.container = append(q.container, x)
}
func (q *Queue) Dequeue() int {
	a := q.container[0]
	q.container = q.container[1:]
	return a
}
func (q *Queue) IsEmpty() bool {
	return len(q.container) == 0
}
