package stack

import _interface "github.com/Johnathan-Chan/dataStructure/pkg/interface"

// IStack stack interface
type IStack[T interface{}, C _interface.Integer] interface {
	Push(T)
	Pop() T
	Capacity() C
	Top() T
}
