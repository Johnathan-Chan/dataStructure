package the_stack_containing_the_min_function

import (
	_interface "github.com/Johnathan-Chan/dataStructure/pkg/interface"
	"github.com/Johnathan-Chan/dataStructure/pkg/stack"
	"github.com/Johnathan-Chan/dataStructure/pkg/stack/linkstack"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type MinStack[T _interface.Number, C _interface.Integer] struct {
	min      T
	stack    stack.IStack[T, C]
	minStack stack.IStack[T, C]
}

func Constructor[T _interface.Number, C _interface.Integer]() MinStack[T, C] {
	return MinStack[T, C]{
		stack:    linkstack.NewLinkStack[T, C](),
		minStack: linkstack.NewLinkStack[T, C](),
	}
}

func (this *MinStack[T, C]) Push(x T) {
	if this.stack.Capacity() == 0 || this.min >= x {
		this.min = x
	}

	this.stack.Push(x)
	this.minStack.Push(this.min)
}

func (this *MinStack[T, C]) Pop() {
	this.stack.Pop()
	this.minStack.Pop()
	this.min = this.Min()
}

func (this *MinStack[T, C]) Top() T {
	return this.stack.Top()
}

func (this *MinStack[T, C]) Min() T {
	return this.minStack.Top()
}
