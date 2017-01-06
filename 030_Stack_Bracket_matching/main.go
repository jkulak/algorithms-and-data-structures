package main

import "github.com/jkulak/go-data-structures"

func main() {

	s := stack.New()
	s = s.Populate([]int{1, 2, 3, 4, 5, 6, 7})
	s = s.Push(4)
	s = s.Push(4)
	s = s.Push(4)
	s, _ = s.Pop()
	s.Print()
}
