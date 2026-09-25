package list

type TWList[T any] struct {
	Data T
	Next *TWList[T]
	Prev *TWList[T]
}

func NewTwList[T any](data T) *TWList[T] {
	return &TWList[T]{
		Data: data,
	}
}

func (l *TWList[T]) Push(data T) {
	node := l

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &TWList[T]{
		Data: data,
		Prev: node,
	}
}

func (l *TWList[T]) PushPop(data T) *TWList[T] {
	var head *TWList[T]
	node := l
	for node != nil && node.Prev != nil {
		node = node.Prev
	}

	head = &TWList[T]{
		Data: data,
		Next: node,
	}

	if node != nil {
		node.Prev = head
	}

	return head
}

func (l *TWList[T]) Reverse() *TWList[T] {
	var head *TWList[T]
	node := l
	for node != nil {
		next := node.Next
		node.Next, node.Prev = node.Prev, node.Next
		head = node
		node = next
	}

	return head
}