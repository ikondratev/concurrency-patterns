package list

type OWList[T any] struct {
	data T
	next *OWList[T]
}

func NewOwList[T any](data T) *OWList[T] {
	return &OWList[T] {
		data: data,
	}
}

func (l *OWList[T]) Data() T {
	return l.data
}

func (l *OWList[T]) Next() Forward[T] {
	if l == nil || l.next == nil {
		return nil
	}
	return l.next
}

func (l *OWList[T]) Push(data T) {
	node := l
	for node.next != nil {
		node = node.next
	}
	node.next = &OWList[T] {
		data: data,
	}
}

func (l *OWList[T]) PushPop(data T) *OWList[T] {
	return &OWList[T]{
		data: data,
		next: l,
	}
}

func (l *OWList[T]) Reverse() *OWList[T] {
	var head *OWList[T]
	node := l
	for node != nil {
		next := node.next
		node.next = head
		head = node
		node = next
	}

	return head
}