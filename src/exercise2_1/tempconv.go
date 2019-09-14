// Package tempconv convert temperature in Celsius and Fahrenheit
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FeezingC      Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%fK", k) }
