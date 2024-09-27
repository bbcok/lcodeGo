package dataStruct

type MyStack struct {
	queue Queue
}

func Constructor() MyStack {
	return MyStack{
		queue: NewQueue(),
	}
}
func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)
}
func (this *MyStack) Size() int {
	return this.queue.Size()
}
func (this *MyStack) Pop() int {
	n := this.Size() - 1
	for n != 0 {
		a := this.queue.Dequeue()
		this.queue.Enqueue(a)
		n--
	}
	return this.queue.Dequeue()
}
