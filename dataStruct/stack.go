package dataStruct

// Stack 定义一个带泛型的栈
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push 添加元素到栈顶
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 从栈顶移除元素并返回
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty 检查栈是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
