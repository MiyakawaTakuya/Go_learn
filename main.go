package main

//パッケージの宣言は必ず１つ mainの時はmainプログラムが書かれていることを指す
//goは静的型付き言語 基本は型を指定して進めるべし
//goは宣言された変数を使わないとエラーになる仕様になっている

import (
	"fmt"
	// "time"
)

// i5 := 500
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// fmt.Println("Hello Go!")
	// fmt.Println(time.Now())

	//明示的な定義
	//var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "100です"
	fmt.Println(s)
	s = "1000です"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int    //0g0あ初期値になっている
	var s3 string //空文字が初期値になっている
	fmt.Println(i3, s3, i3)

	//暗黙的な定義  関数の中でしか定義できない
	//変数名 := 値
	i4 := 400
	fmt.Println(i4)
	i4 = 403 //i4はint型の変数として定義されている
	fmt.Println(i4)

	fmt.Println(i5)

	outer()
	// fmt.Println(s4)  outer関数の中で定義されているので使えない

}
