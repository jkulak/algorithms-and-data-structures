package main

import datastruct "github.com/jkulak/go-data-structures"

func bracketsMatch(s string) bool {
	stack := datastruct.New()

	for _, v := range s {
		if v == 91 || v == 40 {
			stack = stack.Push(v)
		} else {
			if stack.Length() == 0 {
				return false
			}
			if v == 93 && stack.Top() != 91 {
				return false
			}
			if v == 41 && stack.Top() != 40 {
				return false
			}
			stack, _ = stack.Pop()
		}
	}
	return stack.Length() == 0
}
