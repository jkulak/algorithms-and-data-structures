package main

import (
	"container/list"
	"flag"
	"fmt"
	"time"
)

type testData struct {
	a *list.List
	b *list.Element
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

func logTest(key int, val testData, verbose bool) {

	if val.a == val.r {
		fmt.Printf("✅ ")
	} else {
		fmt.Printf("☠️ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.r, val.a, val)
	}
}

/* 2.3 Implement an algorith to delete a node in the middle (i.e., any node
but the first and the last node, not necessarily the exact middle) of a singly
linked list, given only access to that node
*/

func deleteMiddle(e *list.Element) {
	// delete b
	// a -> b -> c -> d
	// a -> c -> d
	next := e.Next()

	if e != nil && next != nil {
		fmt.Printf("%v", e.Value.(int))
		e = next
		fmt.Printf("%v", e.Value.(int))
	}

}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	verbose := *paramVerbose
	testNumber := *paramTestNumber

	start := time.Now()

	l1 := initList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	e1 := l1.Front()
	e1 = e1.Next()
	r1 := initList([]int{1, 3, 4, 5, 6, 7, 8, 9, 10})

	l2 := initList([]int{4, 4, 5, 5, 6, 6, 7, 7, 8, 8})
	e2 := l2.Front()
	e2 = e2.Next()
	e2 = e2.Next()
	r2 := initList([]int{4, 4, 5, 6, 6, 7, 7, 8, 8})

	var tests = []testData{

		testData{l1, e1, r1},
		testData{l2, e2, r2},
	}

	if testNumber > -1 {
		val := tests[testNumber]
		key := testNumber
		deleteMiddle(val.b)
		logTest(key, val, verbose)

		t := l1.Front()
		for t != nil {
			fmt.Printf("%v\n", t.Value.(int))
			t = t.Next()
		}

	} else {
		for key, val := range tests {
			deleteMiddle(val.b)
			logTest(key, val, verbose)
			// fmt.Printf("List: %v", val.a)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
