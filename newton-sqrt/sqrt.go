package main

import (
	"fmt"
	"math"
)

var SqrtNumber float64

func Sqrt(number float64, cost int) float64{
	if cost <= 0 {
		return number
	}
	if SqrtNumber == 0 {
		SqrtNumber = number
	}
	cost--
	return Sqrt(number - (number*number - SqrtNumber)/(2*number), cost)
}

func main() {
	fmt.Println("My func: %q", Sqrt(1237912, 50))
	fmt.Println("Math sqrt: %q", math.Sqrt(1237912))
}
