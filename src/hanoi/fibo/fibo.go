package main

import "fmt"

func fiboItr(n int) int {
	p, q := 0, 1
	for i := 1; i <= n; i++ {
		p, q = q, p+q
	}
	return p
}

func fiboRec(n int) int {
	if n < 2 {
		return n
	}
	return fiboRec(n-2) + fiboRec(n-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, fiboItr: %3d fiboRec: %3d\n", i, fiboItr(i), fiboRec(i))
	}
}
