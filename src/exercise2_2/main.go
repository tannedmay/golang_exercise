package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"exercise2_1/tempconv"
	"exercise2_2/distconv"
	"exercise2_2/wtconv"
)

func main() {

	if len(os.Args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %s\n", err)
			os.Exit(1)
		}
		args := strings.Split(string(b), "\n")
		printMeasurement(args)
	} else {
		printMeasurement(os.Args[1:])
	}
}

func printMeasurement(args []string) {
	for _, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %s\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(val)
		c := tempconv.Celsius(val)
		k := wtconv.Kilogram(val)
		p := wtconv.Pound(val)
		m := distconv.Meter(val)
		ft := distconv.Feet(val)
		fmt.Printf("Temperature: %s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
		fmt.Printf("Weight: %s = %s, %s = %s\n", k, wtconv.KGtoLB(k), p, wtconv.LBtoKG(p))
		fmt.Printf("Distance: %s = %s, %s = %s\n", m, distconv.MtoFT(m), ft, distconv.FTtoM(ft))
	}
}
