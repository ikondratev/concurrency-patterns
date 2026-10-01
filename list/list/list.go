package list

type List[T any] interface {
	Push(T)
	PushPop(T) List[T]
	Reverse() List[T]
}

type Forward[T any] interface {
	Data() T
	Next() Forward[T]
}