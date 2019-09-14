// Package distconv convert weight in ft and pound
package distconv

import (
	"fmt"
)

type Feet float64
type Meter float64

func (d Feet) String() string  { return fmt.Sprintf("%fft", d) }
func (d Meter) String() string { return fmt.Sprintf("%fm", d) }

// FTtoM convert weight from Feet to Meter.
func FTtoM(d Feet) Meter {
	return Meter(d * 3.28084)
}

// MtoFT convert weight from Meter to Feet.
func MtoFT(d Meter) Feet {
	return Feet(d / 3.28084)
}
