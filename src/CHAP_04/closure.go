package main

import f "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	f.Println(nextInt())
	f.Println(nextInt())
	f.Println(nextInt())

	newInt := intSeq()
	f.Println(newInt())
}
