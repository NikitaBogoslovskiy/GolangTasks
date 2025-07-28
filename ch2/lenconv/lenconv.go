package lenconv

import "fmt"

type Meter float64

func (v Meter) String() string {
	return fmt.Sprintf("%.2f m", v)
}

type Feet float64

func (v Feet) String() string {
	return fmt.Sprintf("%.2f ft", v)
}

func MtoF(value Meter) Feet {
	return Feet(value / 0.3048)
}

func FtoM(value Feet) Meter {
	return Meter(value * 0.3048)
}
