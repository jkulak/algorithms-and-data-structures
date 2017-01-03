package main

import "testing"

func Test2(t *testing.T) {
	l := initList([]int{5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test3(t *testing.T) {
	l := initList([]int{5, 5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test4(t *testing.T) {
	l := initList([]int{5, 5, 5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test5(t *testing.T) {
	l := initList([]int{5, 5, 5, 5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test6(t *testing.T) {
	l := initList([]int{5, 5, 5, 6})
	if isPalindrome(l) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test7(t *testing.T) {
	l := initList([]int{5, 5, 6, 6})
	if isPalindrome(l) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test8(t *testing.T) {
	l := initList([]int{5, 6, 6, 6})
	if isPalindrome(l) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test9(t *testing.T) {
	l := initList([]int{5, 6, 6, 5})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test10(t *testing.T) {
	l := initList([]int{1, 5, 6, 6, 5, 1})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test11(t *testing.T) {
	l := initList([]int{1, 5, 6, 8, 6, 5, 1})
	if isPalindrome(l) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test12(t *testing.T) {
	l := initList([]int{1, 5, 6, 8, 6, 5, 1, 0})
	if isPalindrome(l) != false {
		t.Error("expected", false, "got", true)
	}
}
