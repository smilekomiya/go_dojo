package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// メソッドがポイントレシーバーの場合、呼び出し時に変数 or ポインタのいずれかをレシーバーとして取れる。
// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // vは変数でポインタではないが、メソッドでポイントレシーバが自動的に呼び出される。利便性のためv.Scale(2) は (&v).Scale(2)として解釈される。
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
