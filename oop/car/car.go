package car

import (
	"fmt"
)

// Car is a struct for Car
type Car struct {
	remainFuel float64
	position   int
	effiency   float64
	fuelLimit  float64
}

// NewCar is a constructor for Car struct
func NewCar(remain float64, eff float64) *Car {
	nCar := &Car{}
	nCar.remainFuel = remain
	nCar.position = 0
	nCar.effiency = eff
	return nCar
}

// GoFront is a method that called when you want to move car struct forward
// Parameter distance is distance that want to move car
func (c *Car) GoFront(distance int) bool {
	if c.remainFuel-float64(distance) < 0 {
		return false
	}
	c.position += distance
	c.remainFuel -= float64(distance) / c.effiency
	return true
}

// RemainFuel is a getter for variable remainFuel
func (c *Car) RemainFuel() float64 {
	return c.remainFuel
}

// PrintCarInfo is a method that print members of car struct
func (c *Car) PrintCarInfo() {
	/*
		fmt.Println("남은연료:" + strconv.FormatFloat(c.remainFuel, 'g', 2, 64))
	*/
	fmt.Println("남은연료: " + fmt.Sprintf("%.2f", c.remainFuel))
	fmt.Println("현재위치:", c.position)
}
