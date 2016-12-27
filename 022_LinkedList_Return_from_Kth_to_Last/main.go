package main

import (
	"container/list"
	"flag"
	"fmt"
	"time"
)

type testData struct {
	a *list.List
	b int
	r int
}

// Init a list with values from an array for test purposes
func initList(l *list.List, values []int) {
	for _, val := range values {
		l.PushBack(val)
	}
}

func logTest(key int, val testData, res int, verbose bool) {

	if res == val.r {
		fmt.Printf("✅ ")
	} else {
		fmt.Printf("☠️ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.r, res, val)
	}
}

/* 2.2 Implement an algorithm to find the kth to last element of  asingly
linked list
*/

// Wrong implementation, I didn't understand the task
func returnSliceWrong(l *list.List, k int) *list.List {

	e := l.Front()
	for i := 1; i < k; i++ {
		l.Remove(e)
		e = l.Front()
	}
	return l
}

// Wrong implementation, I didn't understand the task
func returnNewSliceWrong(l *list.List, k int) list.List {
	list := list.New()

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		if i > k {
			list.PushBack(e)
		}
	}

	e := l.Front()
	for i := 1; i < k; i++ {
		l.Remove(e)
		e = l.Front()
	}
	return *list

}

// O(n)
func returnKth(l *list.List, k int) int {

	count := 0

	for e := l.Front(); e != nil; e = e.Next() {
		count++
	}
	elemNum := count - k - 1

	e := l.Front()
	for i := 0; i < elemNum; i++ {
		e = e.Next()
	}
	return e.Value.(int)
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	verbose := *paramVerbose
	testNumber := *paramTestNumber

	start := time.Now()

	l1 := list.New()
	initList(l1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	l2 := list.New()
	initList(l2, []int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8})

	var tests = []testData{

		testData{l1, 0, 10},
		testData{l1, 1, 9},
		testData{l1, 2, 8},
		testData{l1, 3, 7},
		testData{l1, 4, 6},
		testData{l1, 5, 5},
		testData{l1, 6, 4},
		testData{l1, 7, 3},
		testData{l1, 8, 2},
		testData{l1, 9, 1},

		testData{l2, 0, 8},
		testData{l2, 1, 8},
		testData{l2, 2, 7},
		testData{l2, 3, 7},
		testData{l2, 4, 6},
		testData{l2, 5, 6},
		testData{l2, 6, 5},
		testData{l2, 7, 5},
		testData{l2, 8, 4},
		testData{l2, 9, 4},
	}

	if testNumber > -1 {
		val := tests[testNumber]
		key := testNumber
		res := returnKth(val.a, val.b)
		logTest(key, val, res, verbose)
	} else {
		for key, val := range tests {
			res := returnKth(val.a, val.b)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
