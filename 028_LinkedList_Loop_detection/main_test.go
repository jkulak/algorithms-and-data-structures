package main

import "testing"

func Test1(t *testing.T) {
	l := initList([]int{5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}
