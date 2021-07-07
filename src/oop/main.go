package main

import (
	"fmt"

	"struct.com/car"
)

func main() {
	nc := car.NewCar(100.0, 10.0)
	hc := car.NewHybridCar(*nc, 100.0)
	/*
		for i := 1; i < 40; i++ {
			if nc.GoFront(i) {
				fmt.Println(i)
				nc.PrintCarInfo()
			}
		}
	*/
	if hc.ChkConvPossible(1.0) != nil {
		fmt.Println(hc.ChkConvPossible(1.0))
	}
}
