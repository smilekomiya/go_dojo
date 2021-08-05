package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println((&v).Abs()) // 当然OK
	fmt.Println(AbsFunc(v))
	// fmt.Println(AbsFunc(&v)) compile error! 関数は特定の型の引数を取る必要がある。

	// メソッドが変数レシーバである場合、呼び出し時に変数 or ポインタいずれかのレシーバとして受け取る事ができる。
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Print(AbsFunc(*p))
}
