package main

import "testing"

func Test1(t *testing.T) {
	t1 := "[]" // Matching
	if bracketsMatch(t1) != true {
		t.Error("expected", true, "got", false)
	}
}

func Test2(t *testing.T) {
	t1 := "[]()(())" // Matching
	if bracketsMatch(t1) != true {
		t.Error("expected", true, "got", false)
	}
}

func TestFirstWrong(t *testing.T) {
	t1 := "][]()(())" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test3FirstMismatched(t *testing.T) {
	t1 := "[[]()(())" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test4(t *testing.T) {
	t1 := "][]()(())" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test5(t *testing.T) {
	t1 := "[]()[(())" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test6(t *testing.T) {
	t1 := "[]()[(()])]" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test7(t *testing.T) {
	t1 := "[]()[([()]))]" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}

func Test8(t *testing.T) {
	t1 := "[]()[([()])(]" // Not matching
	if bracketsMatch(t1) != false {
		t.Error("expected", false, "got", true)
	}
}
