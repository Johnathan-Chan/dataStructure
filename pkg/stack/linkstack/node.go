package linkstack

// INode node interface
type INode[T interface{}] interface {
	GetValue() T
	SetValue(T)
	GetPre() INode[T]
	SetPre(INode[T])
}

// Node link node
type Node[T interface{}] struct {
	val T
	pre INode[T]
}

func NewNode[T interface{}]() *Node[T] {
	return &Node[T]{}
}

func (n *Node[T]) GetValue() T {
	return n.val
}

func (n *Node[T]) SetValue(t T) {
	n.val = t
}

func (n *Node[T]) GetPre() INode[T] {
	return n.pre
}

func (n *Node[T]) SetPre(i INode[T]) {
	n.pre = i
}
