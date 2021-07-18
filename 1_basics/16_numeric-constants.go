package main

import "fmt"

// 数値の定数は、高精度の値。
// 型のない定数はその状況によって必要な方を取る
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big)) Big overflows int
	fmt.Println(needFloat(Big))
}
