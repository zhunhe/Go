package car

import (
	"fmt"
)

type Car struct {
	remainFuel float64
	position   int
	effiency   float64
	fuelLimit  float64
}

func NewCar(remain float64, eff float64) *Car {
	nCar := &Car{}
	nCar.remainFuel = remain
	nCar.position = 0
	nCar.effiency = eff
	return nCar
}

func (c *Car) GoFront(distance int) bool {
	if c.remainFuel-float64(distance) < 0 {
		return false
	}
	c.position += distance
	c.remainFuel -= float64(distance) / c.effiency
	return true
}

func (c *Car) RemainFuel() float64 {
	return c.remainFuel
}

func (c *Car) PrintCarInfo() {
	/*
		fmt.Println("남은연료:" + strconv.FormatFloat(c.remainFuel, 'g', 2, 64))
	*/
	fmt.Println("남은연료: " + fmt.Sprintf("%.2f", c.remainFuel))
	fmt.Println("현재위치:", c.position)
}
