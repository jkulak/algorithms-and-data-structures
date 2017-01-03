package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// Init a list with values from an array for test purposes
func initList(values []int) *list.List {
	l := list.New()
	for _, val := range values {
		l.PushBack(val)
	}
	return l
}

func areListsEqual(a, b *list.List) bool {

	// Same address in memory
	if a == b {
		return true
	}

	if a.Len() != b.Len() {
		return false
	}

	e1 := a.Front()
	e2 := b.Front()
	for e1 != nil && e2 != nil {
		if !reflect.DeepEqual(e1, e2) {
			// if e1.Value.(ini) != e2.Value.(ini) {
			return false
		}
		e1 = e1.Next()
		e2 = e2.Next()
	}
	return true
}

func printList(l *list.List) {
	e := l.Front()
	fmt.Printf("Printing the list at %v\n", &l)
	for e != nil {
		fmt.Println(e.Value.(int))
		e = e.Next()
	}
	fmt.Printf("------\n")
}

/* 2.6 Implement an algorithm to check if a list is a palindrome
 */

func isPalindrome(l *list.List) bool {

	tmp := new(list.List)

	e := l.Front()
	for e != nil {
		tmp.PushFront(e.Value.(int))
		e = e.Next()
	}

	return areListsEqual(l, tmp)
}

func main() {

}
