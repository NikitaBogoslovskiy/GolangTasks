package main

import "fmt"

func main() {
	const (
		boilingF  = 212.0
		freezingF = 32.0
	)
	fmt.Printf("Boiling temperature in F = %g, in C = %g\n", boilingF, ftoc(boilingF))
	fmt.Printf("Freezing temperature in F = %g, in C = %g\n", freezingF, ftoc(freezingF))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
