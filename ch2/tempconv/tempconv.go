package tempconv

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

type Kelvin float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KtoC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}
