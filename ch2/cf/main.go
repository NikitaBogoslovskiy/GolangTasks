package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		c := tempconv.Celsius(v)
		f := tempconv.Fahrenheit(v)
		fmt.Printf(
			"%s = %s, %s = %s\n",
			c, tempconv.CtoF(c),
			f, tempconv.FtoC(f),
		)
	}
}
