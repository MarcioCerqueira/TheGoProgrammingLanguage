package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) //no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name, default value,
// and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
