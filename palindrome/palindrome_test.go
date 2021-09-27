package palindrome

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
	testCase{"12312", false},
}

func TestHasConsonantSuffix(t *testing.T) {
	for _, e := range testCases {
		got := palindrome(e.in)
		if got != e.out {
			t.Errorf("got: [%v], but want: [%v]", got, e.out)
		}
	}
}
