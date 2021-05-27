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

	i, _ = strconv.Atoi("350")                 // 350
	fmt.Sscanf("57", "%d", &i)                 // 57
	k, _ = strconv.ParseInt("cc7fdd", 16, 32)  // 13402077
	k, _ = strconv.ParseInt("0xcc7fdd", 0, 32) // 13402077
	f, _ = strconv.ParseFloat("3.14", 64)      // 3.14
	s = strconv.Itoa(340)                      // "340"
	s = strconv.FormatInt(134020777, 16)       // "cc7fdd"
	s = fmt.Sprint(3.14)                       // "3.14"
	s = fmt.Sprintf("%x", 13402077)            // "cc7fdd"
}
