package hangul

import "testing"

type testCase struct {
	in  string
	out bool
}

var testCases = []testCase{
	testCase{"가", false},
	testCase{"각", true},
}

func TestHasConsonantSuffix(t *testing.T) {
	for _, e := range testCases {
		got := HasConsonantSuffix(e.in)
		if got != e.out {
			t.Errorf("got: [%v], but want: [%v]", got, e.out)
		}
	}
}
