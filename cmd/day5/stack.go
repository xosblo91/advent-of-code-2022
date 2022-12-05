package main

type Stack[T any] struct {
	list []T
}

func (stack *Stack[T]) Peek() T {
	if stack.Empty() {
		panic("Stack is empty")
	}
	return stack.list[len(stack.list)-1]
}

func NewStack[T any](s []T) Stack[T] {
	return Stack[T]{s}
}

func (stack *Stack[T]) Push(elm T) {
	stack.list = append(stack.list, elm)
}

func (stack *Stack[T]) Pop() (elm T) {
	elm, stack.list = stack.Peek(), stack.list[0:len(stack.list)-1]
	return elm
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.list) == 0
}
