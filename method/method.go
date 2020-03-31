package main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

// メソッド(レシーバ値 レシーバ型) 関数名という形。
// Abs()はVertex xのメソッドとなる。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 普通の関数（レシーバなし）
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v:= Vertex{3, 4} // この様にもかける

	fmt.Println(v.Abs()) // レシーバあり
	                     // 構造体Vertexがあるパッケージ内ならこの様に呼出せる。
						 // クラスメソッドみたいな感じか？

	fmt.Println(Abs2(v)) // レシーバなしの関数呼び出し
}
