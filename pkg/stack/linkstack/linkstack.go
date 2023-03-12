// Package linkstack
package linkstack

import _interface "github.com/Johnathan-Chan/dataStructure/pkg/interface"

// LinkStack build stack from link
type LinkStack[T interface{}, C _interface.Integer] struct {
	root     INode[T]
	capacity C
}

// NewLinkStack init the linkstack
func NewLinkStack[T interface{}, C _interface.Integer]() *LinkStack[T, C] {
	return &LinkStack[T, C]{}
}

// Push push val to top of the stack
func (l *LinkStack[T, C]) Push(val T) {
	now := NewNode[T]()
	now.SetValue(val)
	now.SetPre(l.root)
	l.root = now
	l.capacity++
}

// Pop pop val from top of the stack
func (l *LinkStack[T, C]) Pop() T {
	var val T
	if l.root != nil {
		now := l.root
		l.root = l.root.GetPre()
		now.SetPre(nil)
		l.capacity--
		val = now.GetValue()
	}
	return val
}

// Capacity get the capacity of the stack
func (l *LinkStack[T, C]) Capacity() C {
	return l.capacity
}

// Top get top of the stack and not push it
func (l *LinkStack[T, C]) Top() T {
	var val T
	if l.root != nil {
		val = l.root.GetValue()
	}
	return val
}
