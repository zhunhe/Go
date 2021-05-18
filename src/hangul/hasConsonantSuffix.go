package hangul

var (
	start = '가'
	end   = '힣'
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r <= end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}
