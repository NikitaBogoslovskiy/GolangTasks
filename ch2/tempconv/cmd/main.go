package main

import (
	"fmt"

	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CtoK(tempconv.AbsoluteZeroC))
	fmt.Println(tempconv.KtoC(0))
}
