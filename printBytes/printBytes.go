package main

import "fmt"

// result [ea:b0:80:eb:82:98:eb:8b:a4:]
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
}

// result [ea b0 80 eb 82 98 eb 8b a4]
func printBytes2(s string) {
	fmt.Printf("% x\n", s)
}

func main() {
	printBytes("가나다")
	printBytes2("가나다")
}
