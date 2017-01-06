package main

import "testing"

func Test1(t *testing.T) {
	s := newSetOfStacks(4)
	for i := 1; i < 17; i++ {
		s = s.Push('x')
	}
	l := len(s.stacks)
	if l != 4 {
		t.Error("expected", 4, "got", l)
	}
}

func TestEmpty(t *testing.T) {
	s := newSetOfStacks(4)
	l := len(s.stacks)
	if l != 1 {
		t.Error("expected", 1, "got", l)
	}
}

func Test2(t *testing.T) {
	s := newSetOfStacks(4)
	for i := 1; i < 18; i++ {
		s = s.Push('x')
	}
	s, _ = s.Pop()  // New stack still exists - but is at 0 index
	s, _ = s.Pop()  // Removes last one from the previous
	s = s.Push('x') // Adds the last one to the last - still 4 stacks

	l := len(s.stacks)
	if l != 4 {
		t.Error("expected", 4, "got", l)
	}
}
