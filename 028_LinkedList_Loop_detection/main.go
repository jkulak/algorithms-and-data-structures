package main

import (
	"container/list"
	"fmt"
)

// Init a list with values from an array for test purposes
func initList(values []int) *list.List {
	l := list.New()
	for _, val := range values {
		l.PushBack(val)
	}
	return l
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

/* 2.8 Detect if a list is looped
 */

func isLooped(l *list.List) bool {
	res := false

	return res
}

func main() {

}
