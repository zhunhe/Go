package main

import "fmt"

func strCat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	s1, s3, s5 := "abc", "abc", "abc"
	s2, s4, s6 := "def", "def", "def"

	// Case 1
	s1 = strCat(s1, s2)
	// Case 2
	s3 = fmt.Sprint(s3, s4)
	// Case 3
	s5 = fmt.Sprintf("%s%s", s5, s6)

	fmt.Printf("s1: %s, s3: %s, s5: %s\n", s1, s3, s5)
}
