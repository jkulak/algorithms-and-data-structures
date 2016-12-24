package main

import (
	"container/list"
	"flag"
	"fmt"
	"reflect"
	"time"
)

type testData struct {
	a   *list.List
	b   *list.List
	res bool
}

func logTest(key int, val testData, res *list.List, verbose bool) {

	if reflect.DeepEqual(res, val.b) {
		fmt.Printf("✅ ")
	} else {
		fmt.Printf("☠️ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.b, res, val)
	}
}

/* 2.1 Write cod to remove duplicates from an unsorted linked list
2.1a How would you solve this problem if a temp buffer is not allowed?
*/

func removeDups(l *list.List) *list.List {

	dups := make(map[int]int)
	e := l.Front()

	for e != nil {
		// temp next value needed in case we remove the element
		next := e.Next()
		if dups[e.Value.(int)] > 0 {
			l.Remove(e)
		} else {
			dups[e.Value.(int)]++
		}
		e = next
	}
	return l
}

func removeDupsNoBuffer(l *list.List) *list.List {

	e := l.Front()
	for e != nil {
		runner := e.Next()
		for runner != nil {
			runnerNext := runner.Next()
			if e.Value.(int) == runner.Value.(int) {
				l.Remove(runner)
			}
			runner = runnerNext
		}
		e = e.Next()
	}
	return l
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	var verbose bool = *paramVerbose
	var testNumber int = *paramTestNumber

	var start time.Time = time.Now()

	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)
	l1.PushBack(5)
	l1.PushBack(5)

	t1 := list.New()
	t1.PushBack(1)
	t1.PushBack(2)
	t1.PushBack(3)
	t1.PushBack(4)
	t1.PushBack(5)

	l2 := list.New()
	l2.PushBack(5)
	l2.PushBack(2)
	l2.PushBack(5)
	l2.PushBack(4)
	l2.PushBack(5)
	l2.PushBack(5)

	t2 := list.New()
	t2.PushBack(5)
	t2.PushBack(2)
	t2.PushBack(4)

	l3 := list.New()
	l3.PushBack(2)
	l3.PushBack(5)
	l3.PushBack(5)
	l3.PushBack(5)
	l3.PushBack(5)
	l3.PushBack(5)

	t3 := list.New()
	t3.PushBack(2)
	t3.PushBack(5)

	l4 := list.New()
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(5)
	l4.PushBack(2)
	l4.PushBack(3)

	t4 := list.New()
	t4.PushBack(5)
	t4.PushBack(2)
	t4.PushBack(3)

	l5 := list.New()
	l5.PushBack(5)
	l5.PushBack(5)
	l5.PushBack(5)
	l5.PushBack(5)
	l5.PushBack(2)
	l5.PushBack(5)
	l5.PushBack(2)
	l5.PushBack(5)
	l5.PushBack(3)
	l5.PushBack(3)
	l5.PushBack(3)
	l5.PushBack(3)
	l5.PushBack(5)
	l5.PushBack(5)

	t5 := list.New()
	t5.PushBack(5)
	t5.PushBack(2)
	t5.PushBack(3)

	l6 := list.New()
	l6.PushBack(5)
	l6.PushBack(5)

	t6 := list.New()
	t6.PushBack(5)

	l7 := list.New()
	l7.PushBack(5)
	l7.PushBack(5)
	l7.PushBack(1)
	l7.PushBack(1)

	t7 := list.New()
	t7.PushBack(5)
	t7.PushBack(1)

	l8 := list.New()
	l8.PushBack(5)

	t8 := list.New()
	t8.PushBack(5)

	l9 := list.New()

	t9 := list.New()

	var tests = []testData{

		testData{l1, t1, true},
		testData{l2, t2, true},
		testData{l3, t3, true},
		testData{l4, t4, true},
		testData{l5, t5, true},
		testData{l6, t6, true},
		testData{l7, t7, true},
		testData{l8, t8, true},
		testData{l9, t9, true},
	}

	if testNumber > -1 {
		val := tests[testNumber]
		key := testNumber
		res := removeDupsNoBuffer(val.a)
		logTest(key, val, res, verbose)
	} else {
		for key, val := range tests {
			res := removeDupsNoBuffer(val.a)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
