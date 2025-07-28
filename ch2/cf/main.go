package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/lenconv"
	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/tempconv"
	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/weightconv"
)

var temp = flag.Bool("temp", false, "Temperature conversions")
var weight = flag.Bool("weight", false, "Weight conversions")
var len = flag.Bool("len", false, "Length conversions")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}

		fmt.Printf("Conversion results for %g\n", v)

		if *temp {
			fmt.Println("Temperature")
			c := tempconv.Celsius(v)
			f := tempconv.Fahrenheit(v)
			fmt.Printf(
				"%s = %s, %s = %s\n",
				c, tempconv.CtoF(c),
				f, tempconv.FtoC(f),
			)
		}

		if *weight {
			fmt.Println("Weight")
			k := weightconv.Kilogram(v)
			p := weightconv.Pound(v)
			fmt.Printf(
				"%s = %s, %s = %s\n",
				k, weightconv.KtoP(k),
				p, weightconv.PtoK(p),
			)
		}

		if *len {
			fmt.Println("Length")
			m := lenconv.Meter(v)
			f := lenconv.Feet(v)
			fmt.Printf(
				"%s = %s, %s = %s\n",
				m, lenconv.MtoF(m),
				f, lenconv.FtoM(f),
			)
		}

		fmt.Println()
	}
}
