// Write a method to replace all spaces in a string with "%20"
// make it UTF ready
package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func urlify(s string) string {
	result := ""
	str := strings.TrimRight(s, " ")
	len := len(str)
	for i := len - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			result = "%20" + result
		} else {
			result = string(s[i]) + result
		}
	}
	return result
}

func main() {

	var start time.Time
	start = time.Now()

	type testData struct {
		a   string
		res string
	}

	var tests = []testData{
		testData{"a a  ", "a%20a"},
		testData{"b 2  ", "b%202"},
		testData{" 2  ", "%202"},
		testData{"   ", "%20"},
		testData{"      ", "%20%20"},
		testData{"just a test    ", "just%20a%20test"},
	}

	fmt.Printf("%v\n", tests)

	for key, val := range tests {
		res := urlify(val.a)

		if res != val.res {
			fmt.Printf("🚫  ")
		} else {
			fmt.Printf("✅  ")
		}
		fmt.Printf("Test %d %v\n\texpected: \"%s\"\n\treturned: \"%s\"\n", key+1, val, val.res, res)
	}

	log.Printf("Time: %s", time.Since(start))
}
