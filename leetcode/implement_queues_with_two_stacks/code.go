package implement_queues_with_two_stacks

import (
	_interface "github.com/Johnathan-Chan/dataStructure/pkg/interface"
	"github.com/Johnathan-Chan/dataStructure/pkg/stack"
	"github.com/Johnathan-Chan/dataStructure/pkg/stack/linkstack"
)

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

type CQueue[T interface{}, C _interface.Integer] struct {
	left  stack.IStack[T, C]
	right stack.IStack[T, C]
}

func Constructor[T interface{}, C _interface.Integer]() CQueue[T, C] {
	return CQueue[T, C]{
		left:  linkstack.NewLinkStack[T, C](),
		right: linkstack.NewLinkStack[T, C](),
	}
}

func (this *CQueue[T, C]) AppendTail(value T) {
	this.left.Push(value)
}

func (this *CQueue[T, C]) DeleteHead() T {
	if this.right.Capacity() == 0 {
		for this.left.Capacity() != 0 {
			this.right.Push(this.left.Pop())
		}
	}
	return this.right.Pop()
}
