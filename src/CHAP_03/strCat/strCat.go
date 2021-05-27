package main

import "fmt"

func strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)   // result [abcdef]
	fmt.Println(*ps) // result [abcdef]
}

func main() {
	strCat()
}
