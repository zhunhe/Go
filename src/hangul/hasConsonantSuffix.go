package hangul

var (
	start = '가' // rune(44032)
	end   = '힣' // rune(55203)
)

// 받침이 있으면 true, 없으면 false를 return하는 프로그램
// "아무개이(가) 찾아왔습니다."와 같이 매끄럽지 못한 문장이 아니라
// "아무개가 찾아왔습니다", "한글이 찾아왔습니다"와 같은 문장을 사용할 수 있도록..
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
