package main

import "testing"

type testCase struct {
	in  string
	out bool
}

var testCases = []testCase{
	testCase{"1", true},
	testCase{"12321", true},
	testCase{"12345678987654321", true},
	testCase{"12345678977654321", false},
	testCase{"12312", false},
}

func TestCheckPalindrome(t *testing.T) {
	for _, e := range testCases {
		got := CheckPalindrome(e.in)
		if got != e.out {
			t.Errorf("testCase: [%v] got: [%v], but want: [%v]", e.in, got, e.out)
		}
	}
}
