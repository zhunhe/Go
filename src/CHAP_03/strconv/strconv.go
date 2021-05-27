package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var k int64
	var f float64
	var s string

	i, _ = strconv.Atoi("350")
	fmt.Println(i)

	k, _ = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k)

	k, _ = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)

	f, _ = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)

	s = strconv.Itoa(340)
	fmt.Println(s)

	s = strconv.FormatInt(134020777, 16)
	fmt.Println(s)
}
