package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c Celsius) cToF() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%f°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) fToC() Celsius {
	return Celsius(f-32) * 5 / 9
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%f°F", f)
}

type Feet float64

func (f Feet) fToM() Meter {
	return Meter(f * 0.3048)
}

func (f Feet) String() string {
	return fmt.Sprintf("%ff", f)
}

type Meter float64

func (m Meter) mToF() Feet {
	return Feet(m / 0.3048)
}

func (m Meter) String() string {
	return fmt.Sprintf("%fm", m)
}

type Pound float64

func (p Pound) pToK() Kilogram {
	return Kilogram(p * 0.45359237)
}

func (p Pound) String() string {
	return fmt.Sprintf("%flbs", p)
}

type Kilogram float64

func (k Kilogram) kToP() Pound {
	return Pound(k / 0.45359237)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%fKg", k)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	valuePtr := flag.Float64("value", 0.0, "value")
	poundToKgPtr := flag.Bool("ptk", false, "pound to kilogram")
	kgToPound := flag.Bool("ktp", false, "kilogram to pound")
	feetToMeter := flag.Bool("ftm", false, "feet to meter")
	meterToFeet := flag.Bool("mtf", false, "meter to feet")
	cToF := flag.Bool("ctf", false, "celsius to fahrenheit")
	fToC := flag.Bool("ftc", false, "fahrenheit to celsius")

	flag.Parse()

	value := float64(*valuePtr)

	fmt.Println("value:", value)

	switch true {
	case *poundToKgPtr:
		p := Pound(value)
		fmt.Println(p.pToK())
	case *kgToPound:
		k := Kilogram(value)
		fmt.Println(k.kToP())
	case *feetToMeter:
		f := Feet(value)
		fmt.Println(f.fToM())
	case *meterToFeet:
		m := Meter(value)
		fmt.Println(m.mToF())
	case *cToF:
		c := Celsius(value)
		fmt.Println(c.cToF())
	case *fToC:
		f := Fahrenheit(value)
		fmt.Println(f.fToC())
	}
}
