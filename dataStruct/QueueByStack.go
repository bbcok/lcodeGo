package dataStruct

type MyQueue struct {
	stack1 Stack[int]
	stack2 Stack[int]
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		stack1: Stack[int]{
			items: make([]int, 0),
		},
		stack2: Stack[int]{
			items: make([]int, 0),
		},
	}
}

func (q *MyQueue) Enqueue(item int) {
	q.stack1.Push(item)
}
func (q *MyQueue) Dequeue() int {
	for !q.stack1.IsEmpty() {
		a, _ := q.stack1.Pop()
		q.stack2.Push(a)
	}
	ans, _ := q.stack2.Pop()
	for !q.stack2.IsEmpty() {
		a, _ := q.stack2.Pop()
		q.stack1.Push(a)
	}
	return ans
}
