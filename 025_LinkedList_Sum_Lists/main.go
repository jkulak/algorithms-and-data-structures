package main

import (
	"container/list"
	"flag"
	"fmt"
	"math"
	"reflect"
	"time"
)

type testData struct {
	a *list.List
	b *list.List
	r *list.List
}

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

func logTest(key int, val testData, res list.List, verbose bool) {

	if areListsEqual(val.r, &res) {
		fmt.Printf("✅ ")
	} else {
		fmt.Printf("☠️ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.r, val.a, val)
	}
}

/* 2.5 You have two numbers represented by a linked list, where each node
contains a singl digit. The digits are stored in reverse order, sich that 1's
digit is at the head of the list. Write a function that adds the two numbers and
returns the sum as a linkd list.
*/

func sumLists(a, b *list.List) list.List {
	sum := listValue(a) + listValue(b)
	res := new(list.List)
	for sum > 0 {
		res.PushBack(sum % 10)
		sum = sum / 10
	}
	return *res
}

func listValue(a *list.List) int {
	val := 0
	pow := 0
	e := a.Front()
	for e != nil {
		val = val + e.Value.(int)*int(math.Pow10(pow))
		pow++
		e = e.Next()
	}
	return val
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	verbose := *paramVerbose
	testNumber := *paramTestNumber

	start := time.Now()

	a1 := initList([]int{1, 1, 1})
	b1 := initList([]int{2, 2, 2})
	r1 := initList([]int{3, 3, 3})

	a2 := initList([]int{0, 1, 1})
	b2 := initList([]int{2, 2, 2})
	r2 := initList([]int{2, 3, 3})

	a3 := initList([]int{0, 1, 1})
	b3 := initList([]int{2, 2, 2})
	r3 := initList([]int{2, 3, 3})

	a4 := initList([]int{0, 1, 1, 1})
	b4 := initList([]int{2, 2, 2})
	r4 := initList([]int{2, 3, 3, 1})

	a5 := initList([]int{2, 1, 0})
	b5 := initList([]int{5, 4, 3})
	r5 := initList([]int{7, 5, 3})

	a6 := initList([]int{2, 6, 0})
	b6 := initList([]int{9, 4, 3})
	r6 := initList([]int{1, 1, 4})

	a7 := initList([]int{2, 6, 0})
	b7 := initList([]int{9, 9, 9})
	r7 := initList([]int{1, 6, 0, 1})

	var tests = []testData{

		testData{a1, b1, r1},
		testData{a2, b2, r2},
		testData{a3, b3, r3},
		testData{a4, b4, r4},
		testData{a5, b5, r5},
		testData{a6, b6, r6},
		testData{a7, b7, r7},
	}

	if testNumber > -1 {
		val := tests[testNumber]
		res := sumLists(val.a, val.b)
		logTest(testNumber, val, res, verbose)
	} else {
		for key, val := range tests {
			res := sumLists(val.a, val.b)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
