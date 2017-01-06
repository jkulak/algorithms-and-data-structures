package main

import datastruct "github.com/jkulak/go-data-structures"

type setOfStacks struct {
	stacks       []datastruct.Stack
	currStack    int
	maxStackSize int
}

func (s setOfStacks) Push(v rune) setOfStacks {
	if s.stacks[s.currStack].Length() == s.maxStackSize {
		s.stacks = append(s.stacks, datastruct.NewStack())
		s.currStack++
	}
	s.stacks[s.currStack] = s.stacks[s.currStack].Push(v)
	return s
}

func (s setOfStacks) Pop() (setOfStacks, rune) {
	// fmt.Printf("stacks: %d, current: %d, currentLen: %d\n", len(s.stacks), s.currStack, s.stacks[s.currStack].Length())
	if s.stacks[s.currStack].Length() == 0 {
		i := len(s.stacks)

		s.stacks = append(s.stacks[:i-1])
		s.currStack--
	}
	tmp, elem := s.stacks[s.currStack].Pop()
	s.stacks[s.currStack] = tmp
	return s, elem
}

// newSetOfStacks creates and returns a setOfStacks
func newSetOfStacks(maxStackSize int) setOfStacks {
	stacks := []datastruct.Stack{datastruct.NewStack()}
	return setOfStacks{stacks, 0, maxStackSize}
}
