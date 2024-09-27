package main

type MonotonousQueue struct {
	data []int
}

func NewMonotonousQueue() *MonotonousQueue {
	return &MonotonousQueue{
		data: make([]int, 0),
	}
}
func (q *MonotonousQueue) Back() int {
	return q.data[len(q.data)-1]
}
func (q *MonotonousQueue) Front() int {
	return q.data[0]
}
func (q *MonotonousQueue) Size() int {
	return len(q.data)
}
func (q *MonotonousQueue) Enqueue(el int) {
	for q.Size() > 0 && el > q.Back() {
		q.data = q.data[:len(q.data)-1]
	}
	q.data = append(q.data, el)
}
func (q *MonotonousQueue) Dequeue(el int) {
	if q.Size() > 0 && el == q.data[0] {
		q.data = q.data[1:]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	queue := NewMonotonousQueue()
	for i := 0; i < k; i++ {
		queue.Enqueue(nums[i])
	}
	ans = append(ans, queue.Front())
	for i := k; i < len(nums); i++ {
		queue.Dequeue(nums[i-k])
		queue.Enqueue(nums[i])
		ans = append(ans, queue.Front())
	}
	return ans
}
