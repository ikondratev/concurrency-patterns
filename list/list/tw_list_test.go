package list

import (
	"slices"
	"testing"
)

func valuesTW(l *TWList[int]) []int {
	var got []int
	for n := l; n != nil; n = n.next {
		got = append(got, n.data)
	}
	return got
}

func assertTwList(t *testing.T, l *TWList[int], want []int) {
	t.Helper()

	got := valuesTW(l)
	if !slices.Equal(got, want) {
		t.Fatalf("forward: got %v, want %v", got, want)
	}

	var prev *TWList[int]
	for n := l; n != nil; n = n.next {
		if n.prev != prev {
			t.Fatalf("node %d: Prev points to the wrong node", n.data)
		}
		prev = n
	}
}

func TestPushTW(t *testing.T) {
	l := NewTwList(3)
	l.Push(5)
	l.Push(7)

	assertTwList(t, l, []int{3, 5, 7})
}

func TestPushPopTW(t *testing.T) {
	l := NewTwList(5)
	l.Push(7)
	l = l.PushPop(3)

	assertTwList(t, l, []int{3, 5, 7})
}

func TestPushPopTWNil(t *testing.T) {
	var l *TWList[int]
	l = l.PushPop(1)

	assertTwList(t, l, []int{1})
}

func TestReverseTW(t *testing.T) {
	l := NewTwList(5)
	l.Push(7)
	l.Push(9)

	l = l.Reverse()
	assertTwList(t, l, []int{9, 7, 5})
}

func TestReverseTWSingle(t *testing.T) {
	l := NewTwList(5).Reverse()
	assertTwList(t, l, []int{5})
}

func TestReverseTWNil(t *testing.T) {
	var l *TWList[int]
	if got := l.Reverse(); got != nil {
		t.Fatalf("got %#v, want nil", got)
	}
}