// Implement an algorithm to determine if a string has all unique runesPresent
// make it UTF ready
package main

import (
	"log"
	"time"
)

const (
	n = 90
)

// Not using extra data structures
func isUnique(s string) bool {
	max := len(s)
	for a := 0; a < max; a++ {
		for b := 1; b < max-a; b++ {
			if s[a] == s[b+a] {
				return false
			}
		}
	}
	return true
}

// UTF8 safe version of isUnique
// Using an extra data structure map with rune->present
func isUniqueUtf(s string) bool {
	runeString := []rune(s)
	max := len(runeString)
	result := true
	runesPresent := make(map[rune]bool)

	for a := 0; a < max; a++ {
		if runesPresent[runeString[a]] == true {
			return false
		}
		runesPresent[runeString[a]] = true
	}
	return result
}

func main() {

	var start time.Time
	start = time.Now()

	log.Printf("%v", isUniqueUtf("s") == true)
	log.Printf("%v", isUniqueUtf("ss") == false)
	log.Printf("%v", isUniqueUtf("sa") == true)
	log.Printf("%v", isUniqueUtf("sas") == false)
	log.Printf("%v", isUniqueUtf("saa") == false)
	log.Printf("%v", isUniqueUtf("saas") == false)
	log.Printf("%v", isUniqueUtf("saaa") == false)
	log.Printf("%v", isUniqueUtf("abcd") == true)
	log.Printf("%v", isUniqueUtf("abcdd") == false)
	log.Printf("%v", isUniqueUtf("ab👾cda") == false)
	log.Printf("%v", isUniqueUtf("abcdb") == false)
	log.Printf("%v", isUniqueUtf("abcdc") == false)
	log.Printf("%v", isUniqueUtf("abcde") == true)
	log.Printf("%v", isUniqueUtf("abcdefghijklmn") == true)
	log.Printf("%v", isUniqueUtf("abcdefghijklmna") == false)
	log.Printf("%v", isUniqueUtf("abcdefgaijklmna") == false)
	log.Printf("%v", isUniqueUtf("abcdefgaijalmna") == false)
	log.Printf("%v", isUniqueUtf("abcdefgaijalmna") == false)
	log.Printf("%v", isUniqueUtf("aaaaaaaaaaaaaaaaaaaaaaa") == false)
	log.Printf("%v", isUniqueUtf("aaaaaaaaaaacaaabaaaaaaa") == false)
	log.Printf("%v", isUniqueUtf("abcdefghijk") == true)
	log.Printf("%v", isUniqueUtf("abcdefghijk-ł") == true)
	log.Printf("%v", isUniqueUtf("😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀") == false)
	log.Printf("%v", isUniqueUtf("‚„Ļĺłś…Ķ≤Ń™€ßį§¶•Ľľ😀💻🐳🎶🐡🐤😙ąabcdefghijkć√-") == true)

	log.Printf("%v", isUnique("s") == true)
	log.Printf("%v", isUnique("ss") == false)
	log.Printf("%v", isUnique("sa") == true)
	log.Printf("%v", isUnique("sas") == false)
	log.Printf("%v", isUnique("saa") == false)
	log.Printf("%v", isUnique("saas") == false)
	log.Printf("%v", isUnique("saaa") == false)
	log.Printf("%v", isUnique("abcd") == true)
	log.Printf("%v", isUnique("abcdd") == false)
	log.Printf("%v", isUnique("abcda") == false)
	log.Printf("%v", isUnique("abcdb") == false)
	log.Printf("%v", isUnique("abcdc") == false)
	log.Printf("%v", isUnique("abcde") == true)
	log.Printf("%v", isUnique("abcdefghijklmn") == true)
	log.Printf("%v", isUnique("abcdefghijklmna") == false)
	log.Printf("%v", isUnique("abcdefgaijklmna") == false)
	log.Printf("%v", isUnique("abcdefgaijalmna") == false)
	log.Printf("%v", isUnique("abcdefgaijalmna") == false)
	log.Printf("%v", isUnique("aaaaaaaaaaaaaaaaaaaaaaa") == false)
	log.Printf("%v", isUnique("aaaaaaaaaaacaaabaaaaaaa") == false)
	log.Printf("%v", isUnique("abcdefghijk") == true)
	log.Printf("%v", isUnique("abcdefghijk-ł") == true)

	log.Printf("Time: %s", time.Since(start))
}
