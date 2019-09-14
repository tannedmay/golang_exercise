// Package wtconv convert weight in kg and pound
package wtconv

import (
	"fmt"
)

type Kilogram float64
type Pound float64

func (w Kilogram) String() string { return fmt.Sprintf("%fkg", w) }
func (w Pound) String() string    { return fmt.Sprintf("%flb", w) }

// KGtoLB convert weight from Kilogram to Pound.
func KGtoLB(w Kilogram) Pound {
	return Pound(w * 2.204623)
}

// LBtoKG convert weight from Pound to Kilogram.
func LBtoKG(w Pound) Kilogram {
	return Kilogram(w / 2.204623)
}
