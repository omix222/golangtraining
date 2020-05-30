package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 戻り値の型は）の後ろに書く
func add(x int, y int) int {
	return x + y
}

// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して書ける
func add2(x, y int) int {
	return x + y
}

// Multiple results　複数戻り値
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, aa世界")
	fmt.Println("The time is", time.Now())
	//擬似乱数を返す rand.Intn はいつも同じ数を返す
	fmt.Println("My favarite number is ", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println("1+1= ", add(1, 1))
	fmt.Println("1+1= ", add2(1, 1))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
