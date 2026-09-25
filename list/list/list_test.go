package list

import (
	"slices"
	"testing"
)

func values(l *TWList[int]) []int {
	var got []int
	for n := l; n != nil; n = n.Next {
		got = append(got, n.Data)
	}
	return got
}

func assertList(t *testing.T, l *TWList[int], want []int) {
	t.Helper()

	got := values(l)
	if !slices.Equal(got, want) {
		t.Fatalf("forward: got %v, want %v", got, want)
	}

	var prev *TWList[int]
	for n := l; n != nil; n = n.Next {
		if n.Prev != prev {
			t.Fatalf("node %d: Prev points to the wrong node", n.Data)
		}
		prev = n
	}
}

func TestPush(t *testing.T) {
	l := NewTwList(3)
	l.Push(5)
	l.Push(7)

	assertList(t, l, []int{3, 5, 7})
}

func TestPushPop(t *testing.T) {
	l := NewTwList(5)
	l.Push(7)
	l = l.PushPop(3)

	assertList(t, l, []int{3, 5, 7})
}

func TestPushPopNil(t *testing.T) {
	var l *TWList[int]
	l = l.PushPop(1)

	assertList(t, l, []int{1})
}

func TestReverse(t *testing.T) {
	l := NewTwList(5)
	l.Push(7)
	l.Push(9)

	l = l.Reverse()
	assertList(t, l, []int{9, 7, 5})
}

func TestReverseSingle(t *testing.T) {
	l := NewTwList(5).Reverse()
	assertList(t, l, []int{5})
}

func TestReverseNil(t *testing.T) {
	var l *TWList[int]
	if got := l.Reverse(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}