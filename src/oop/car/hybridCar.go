package car

import "errors"

var (
	ErrBatteryFull   = errors.New("car: 배터리가 많아요")
	ErrNotEnoughFuel = errors.New("car: 연료가 없어요")
)

type HybridCar struct {
	Car
	remainBattery     float64
	batteryLimit      float64
	batteryEfficiency float64
	batteryPerFuel    float64
}

func NewHybridCar(c Car, elec float64) *HybridCar {
	hc := &HybridCar{}
	hc.remainBattery = elec
	hc.remainFuel = c.remainFuel
	hc.position = c.position
	hc.batteryEfficiency = 20
	hc.batteryPerFuel = 3.0
	return hc
}

func (hc *HybridCar) calBatteryPerFuel(fuel float64) float64 {
	return hc.batteryPerFuel * fuel
}

func (hc *HybridCar) ChkConvPossible(fuel float64) error {
	if hc.remainFuel-fuel < 0 {
		return ErrNotEnoughFuel
	}
	if hc.remainBattery+hc.calBatteryPerFuel(fuel) > hc.batteryLimit {
		return ErrBatteryFull
	}
	return nil
}

func (hc *HybridCar) ConvFuelToElec(fuel float64) bool {
	hc.remainBattery += hc.calBatteryPerFuel(fuel)
	hc.remainFuel -= fuel
	return true
}
