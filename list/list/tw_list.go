package list

type TWList[T any] struct {
	data T
	next *TWList[T]
	prev *TWList[T]
}

func NewTwList[T any](data T) *TWList[T] {
	return &TWList[T]{
		data: data,
	}
}

func (l *TWList[T]) Data() T {
	return l.data
}

func (l *TWList[T]) Next() Forward[T] {
	if l == nil || l.next == nil {
		return nil
	}

	return l.next
}

func (l *TWList[T]) Push(data T) {
	node := l
	for node.next != nil {
		node = node.next
	}

	node.next = &TWList[T]{
		data: data,
		prev: node,
	}
} 

func (l *TWList[T]) PushPop(data T) *TWList[T] {
	var head *TWList[T]
	node := l
	for node != nil && node.prev != nil {
		node = node.prev
	}

	head = &TWList[T]{
		data: data,
		next: node,
	}
	if node != nil {
		node.prev = head
	}

	return head
}

func (l *TWList[T]) Reverse() *TWList[T] {
	var head *TWList[T]
	node := l
	for node != nil {
		next := node.next
		node.next, node.prev = node.prev, node.next
		head = node
		node = next
	}

	return head
}

