package main

import (
	"fmt"

	"github.com/zhunhe/Go/src/CHAP_03/hangul"
)

func main() {
	fruits := [...]string{"사과", "바나나", "자몽"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
}
