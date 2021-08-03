package main

import "fmt"

type month int

const (
	January = 1 + iota
	Feburay
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
	monthMax
)

func main() {
	//for i := month(1); i < monthMax; i++ {
	for i := January; i < monthMax; i++ {
		fmt.Println(i)
	}
}
