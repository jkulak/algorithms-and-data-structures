package main

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

type testData struct {
	a   Matrix
	b   Matrix
	res bool
}

func logTest(key int, val testData, res Matrix, verbose bool) {
	if reflect.DeepEqual(res, val.b) {
		fmt.Printf("☠️ ")
	} else {
		fmt.Printf("✅ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.b, res, val)
	}
}

/* 1.7 Given an image represented by an NxN matrix, where each pixel in the image
is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this
in place? https://en.wikipedia.org/wiki/In-place_algorithm
*/

// Px type represents a pixel with 4 bytes
type Px [4]byte

// Matrix type is a 2D array of Pixels
type Matrix [][]Px

func rotatMatrix90(a Matrix) Matrix {

	// ...
	return a
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	var verbose bool = *paramVerbose
	var testNumber int = *paramTestNumber

	var start time.Time = time.Now()

	// m[row][column]

	// 0,0 -> 0,2  0,1 -> 1,2  0,2 -> 2,2
	// 1,0 -> 0,1  1,1 -> 1,1  1,2 -> 2,1
	// 2,0 -> 0,0  2,1 -> 1,0  2,2 -> 2,0

	m := Matrix{
		[]Px{Px{1, 2, 3, 1}, Px{1, 2, 3, 2}, Px{1, 2, 3, 3}},
		[]Px{Px{1, 2, 3, 4}, Px{1, 2, 3, 5}, Px{1, 2, 3, 6}},
		[]Px{Px{1, 2, 3, 7}, Px{1, 2, 3, 8}, Px{1, 2, 3, 9}},
	}

	r := Matrix{
		[]Px{Px{1, 2, 3, 7}, Px{1, 2, 3, 4}, Px{1, 2, 3, 1}},
		[]Px{Px{1, 2, 3, 8}, Px{1, 2, 3, 5}, Px{1, 2, 3, 2}},
		[]Px{Px{1, 2, 3, 9}, Px{1, 2, 3, 6}, Px{1, 2, 3, 3}},
	}

	var tests = []testData{

		testData{m, r, true},
	}

	if testNumber > 0 {
		val := tests[testNumber]
		key := testNumber
		res := rotatMatrix90(val.a)
		logTest(key, val, res, verbose)
	} else {
		for key, val := range tests {
			res := rotatMatrix90(val.a)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
