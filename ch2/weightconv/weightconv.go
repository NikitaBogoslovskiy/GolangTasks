package weightconv

import "fmt"

type Kilogram float64

func (v Kilogram) String() string {
	return fmt.Sprintf("%.2f kg", v)
}

type Pound float64

func (v Pound) String() string {
	return fmt.Sprintf("%.2f lb", v)
}

func KtoP(value Kilogram) Pound {
	return Pound(2.205 * value)
}

func PtoK(value Pound) Kilogram {
	return Kilogram(value / 2.205)
}
