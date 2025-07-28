package main

import (
	"fmt"
	"time"

	"github.com/NikitaBogoslovskiy/GolangTasks/ch2/popcount"
)

func main() {
	// var x uint64 = 510
	// fmt.Printf("%d\n", uint8(x))

	var num uint64 = 2784353547892345789
	testTime(func() { fmt.Printf("%d\n", popcount.PopCount1(num)) })
	testTime(func() { fmt.Printf("%d\n", popcount.PopCount2(num)) })
	testTime(func() { fmt.Printf("%d\n", popcount.PopCount3(num)) })
}

func testTime(f func()) {
	var t = time.Now()
	f()
	fmt.Printf("Time: %d\n", time.Since(t).Nanoseconds())
}
