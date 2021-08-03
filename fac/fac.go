package main

import "fmt"

// fac returns n!
func fac(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fac(n-1)
}

// facItr returns n!
func facItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

// facItr2 returns n!
func facItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Printf("fac(5): %d\n", fac(5))
	fmt.Printf("facItr(5): %d\n", facItr(5))
	fmt.Printf("facItr2(5): %d\n", facItr2(5))
}
