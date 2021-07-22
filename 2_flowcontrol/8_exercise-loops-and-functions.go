package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtWithCount(x float64, count int) float64 {
	z := 1.0
	for i := 0; i < count; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtWithLimit(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		var beforeZ = z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-beforeZ)/z < math.Pow(0.1, 10) {
			fmt.Printf("%d 回目で目標の誤差に到達\n", i)
			return z
		}
	}
	fmt.Printf("100回では近似せず\n")
	return z
}

func main() {
	// fmt.Println(
	// 	math.Sqrt(1),
	// 	SqrtWithCount(1, 1),
	// 	SqrtWithCount(1, 5),
	// 	SqrtWithCount(1, 10),
	// 	SqrtWithCount(1, 20),
	// 	SqrtWithCount(1, 100),
	// )

	// fmt.Println(
	// 	math.Sqrt(2),
	// 	SqrtWithCount(2, 1),
	// 	SqrtWithCount(2, 5),
	// 	SqrtWithCount(2, 10),
	// 	SqrtWithCount(2, 20),
	// 	SqrtWithCount(2, 100),
	// )

	// fmt.Println(
	// 	math.Sqrt(3),
	// 	SqrtWithCount(3, 1),
	// 	SqrtWithCount(3, 5),
	// 	SqrtWithCount(3, 10),
	// 	SqrtWithCount(3, 20),
	// 	SqrtWithCount(3, 100),
	// )

	fmt.Println(
		SqrtWithLimit(1),
		SqrtWithLimit(2),
		SqrtWithLimit(3),
		SqrtWithLimit(10),
		SqrtWithLimit(17),
		SqrtWithLimit(21),
		SqrtWithLimit(99),
		SqrtWithLimit(1000000),
	)
}
