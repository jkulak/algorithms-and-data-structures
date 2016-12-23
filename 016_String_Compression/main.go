package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
	"time"
)

type testData struct {
	a   string
	b   string
	res bool
}

func logTest(key int, val testData, res string, verbose bool) {
	if res != val.b {
		fmt.Printf("☠️ ")
	} else {
		fmt.Printf("✅ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.b, res, val)
	}
}

/* 1.6 Implement a method to perform basic string compression using the counts
of repeates characters. For example, the string aabcccccaaa would become a2b1c5a3.
If the "compressed" string would not become smaller than the original string,
your method should return the original string. You can assume the string has only
uppercase and lowercase letters (a-z).
*/

// O(n)
func compressString(a string) string {
	var result bytes.Buffer
	lena := len(a)
	lastChar := a[0:1]
	charCount := 0

	for i := 0; i < lena; i++ {
		if string(a[i]) != lastChar {
			result.WriteString(lastChar + strconv.Itoa(charCount))
			charCount = 0
		}
		lastChar = string(a[i])
		charCount++
	}
	result.WriteString(lastChar + strconv.Itoa(charCount))

	if len(result.String()) >= lena {
		return a
	}
	return result.String()
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	var verbose bool = *paramVerbose
	var testNumber int = *paramTestNumber

	var start time.Time = time.Now()

	var tests = []testData{

		testData{"a", "a", true},
		testData{"aa", "aa", true},
		testData{"aaa", "a3", true},
		testData{"ab", "ab", true},
		testData{"abb", "abb", true},
		testData{"aab", "aab", true},
		testData{"aaab", "aaab", true},
		testData{"aaabb", "a3b2", true},
		testData{"aaabbb", "a3b3", true},
		testData{"aabb", "aabb", true},
		testData{"aabbc", "aabbc", true},
		testData{"aabbcc", "aabbcc", true},
		testData{"aabbbcc", "a2b3c2", true},
		testData{"aaabbbcc", "a3b3c2", true},
		testData{"aaabbbccc", "a3b3c3", true},
		testData{"aaabbbcccc", "a3b3c4", true},
		testData{"cccbbbccccc", "c3b3c5", true},
		testData{"aaabbbcccccd", "a3b3c5d1", true},
		testData{"abcdefffffff", "abcdefffffff", true},
		testData{"fffffffabcde", "fffffffabcde", true},
		testData{"abcdeffffffff", "a1b1c1d1e1f8", true},
		testData{"abcdefffffffff", "a1b1c1d1e1f9", true},
		testData{"abcdefggggggggggh", "abcdefggggggggggh", true},
		testData{"ggggggggggabcdefh", "ggggggggggabcdefh", true},
		testData{"abcdefhgggggggggg", "abcdefhgggggggggg", true},
		testData{"abcdefgggggggggggh", "a1b1c1d1e1f1g11h1", true},
		testData{"abcdefhggggggggggg", "a1b1c1d1e1f1h1g11", true},
		testData{"gggggggggggabcdefh", "g11a1b1c1d1e1f1h1", true},
		testData{"aabbccddeee", "a2b2c2d2e3", true},
		testData{"aabcccccaaa", "a2b1c5a3", true},
	}

	if testNumber > 0 {
		val := tests[testNumber]
		key := testNumber
		res := compressString(val.a)
		logTest(key, val, res, verbose)
	} else {
		for key, val := range tests {
			res := compressString(val.a)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
