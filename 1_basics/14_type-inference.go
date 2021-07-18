package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	x := 42
	fmt.Printf("x is of type %T\n", x)

	y := 42.0
	fmt.Printf("y is of type %T\n", y)

	z := math.Sqrt(7)
	fmt.Printf("z is of type %T\n", z)

	w := cmplx.Sqrt(7)
	fmt.Printf("w is of type %T\n", w)
}
