package hangul

import "testing"

type consonantSuffixTest struct {
	in  string
	out bool
}

var consonantSuffixTests = []consonantSuffixTest{
	consonantSuffixTest{"가", false},
	consonantSuffixTest{"각", true},
}

func TestHasConsonantSuffix(t *testing.T) {
	for _, e := range consonantSuffixTests {
		got := HasConsonantSuffix(e.in)
		if got != e.out {
			t.Errorf("got: [%v], but want: [%v]", got, e.out)
		}
	}
}
