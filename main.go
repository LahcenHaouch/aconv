package main

import "fmt"

type Celsius float64

func (c Celsius) cToF() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%v°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) fToC() Celsius {
	return Celsius(f-32) * 5 / 9
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%v°F", f)
}

type Feet float64

func (f Feet) fToM() Meter {
	return Meter(f * 0.3048)
}

func (f Feet) String() string {
	return fmt.Sprintf("%vf", f)
}

type Meter float64

func (m Meter) mToF() Feet {
	return Feet(m / 0.3048)
}

func (m Meter) String() string {
	return fmt.Sprintf("%vm", m)
}

type Pound float64

func (p Pound) pToK() Kilogram {
	return Kilogram(p * 0.45359237)
}

func (p Pound) String() string {
	return fmt.Sprintf("%vlbs", p)
}

type Kilogram float64

func (k Kilogram) kToP() Pound {
	return Pound(k / 0.45359237)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%vKg", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
}
