package list

import (
	"slices"
	"testing"
)

func valuesOw(l *OWList[int]) []int {
	var got []int
	for n := l; n != nil; n = n.next {
		got = append(got, n.data)
	}
	return got
}

func assertListOW(t *testing.T, l *OWList[int], want []int) {
	t.Helper()

	got := valuesOw(l)
	if !slices.Equal(got, want) {
		t.Fatalf("forward: got %v, want %v", got, want)
	}
}

func TestPushOW(t *testing.T) {
	l := NewOwList(3)
	l.Push(5)
	l.Push(7)

	assertListOW(t, l, []int{3, 5, 7})
}

func TestPushPop(t *testing.T) {
	l := NewOwList(5)
	l.Push(7)
	l = l.PushPop(3)

	assertListOW(t, l, []int{3, 5, 7})
}

func TestPushPopNil(t *testing.T) {
	var l *OWList[int]
	l = l.PushPop(1)

	assertListOW(t, l, []int{1})
}

func TestReverse(t *testing.T) {
	l := NewOwList(5)
	l.Push(7)
	l.Push(9)

	l = l.Reverse()
	assertListOW(t, l, []int{9, 7, 5})
}

func TestReverseSingle(t *testing.T) {
	l := NewOwList(5).Reverse()
	assertListOW(t, l, []int{5})
}

func TestReverseNil(t *testing.T) {
	var l *OWList[int]
	if got := l.Reverse(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}