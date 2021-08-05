package main

import (
	"fmt"
	"math"
)

// ポインタレシーバを使う２つの理由
// - メソッドがレシーバが指す先の変数を変更する。
// - メソッドの呼び出しごとに変数のコピーを避ける。
// 	- レシーバが大きな構造体である場合に有効

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ポインタレシーバである必要はないがそうなっている。
// 一般的には値レシーバ or ポインタレシーバのどちらかですべてのメソッドを与えるべき（後述）
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	// Printf は %+v のようにプラスを与えることでフィールド名も表示される。
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
