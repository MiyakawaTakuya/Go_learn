package main

//パッケージの宣言は必ず１つ mainの時はmainプログラムが書かれていることを指す
//goは静的型付き言語 基本は型を指定して進めるべし
//goは宣言された変数を使わないとエラーになる仕様になっている

import (
	"fmt"
	"strconv"
	// "time"
)

// i5 := 500
var i5 int = 500

const Pi = 3.14
const (
	URL      = "https://xxx.co.jp"
	SiteName = "test"
)
const ( //1 1 1 A A A
	A = 1
	B
	C
	D = "A"
	E
	F
)
const (
	c0 = iota //連番を返す
	c1
	c2
)

// var Big int = 9223372036854775807
// var Big int = 9223372036854775807+1 //これはエラーになる
const Big = 9223372036854775807 + 1 //定数ではエラーにならない

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func Plus(x int, y int) int {
	return x + y + y
}

func Plus_(x, y int) int { //型の指定は片方に省略することもできる
	return x + y + y
}

func Div(x, y int) (int, int) { //型の指定は片方に省略することもできる
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return //返り値はない場合
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func CallFunc(f func()) {
	f()
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {
	////function
	fmt.Println(Plus(2, 4))
	sho, amari := Div(33, 5)
	fmt.Println(sho, amari)
	fmt.Println(Div(33, 5))
	sho_, _ := Div(33, 6) //_にすることで返り値を使いたくないときは破棄できる
	fmt.Println(sho_)
	fmt.Println(Double(6))
	Noreturn()
	CallFunc(func() {
		fmt.Println("Function as para")
	})

	////anonymous function
	f1 := func(x, y int) int {
		return x + y
	}
	i26 := f1(543, 12)
	fmt.Println(i26)

	f2 := func(x, y int) int {
		return x + y
	}(20, 2) //無名関数の省略した書き方
	fmt.Println(f2) //22

	////function return function
	f3 := ReturnFunc() //関数をf3無名関数として代入(？)している
	f3()               //I'm a function

	////Goの無名関数はクロージャーで、クロージャーとは日本語では関数閉包と呼ばれ、
	//関数と関数の処理に関する関数外の環境をセットして閉じ込めた物
	//クロージャーの中で定義された変数は呼び出されるたびに初期化されず、値が生き続ける
	//TODO クロージャー復習

	////ジェネレーター  何らかのルールに従って連続した値を返し続ける仕組みの事
	//クロージャーを応用して、ジェネレーターを実装する
	ints :=integers()
	fmt.Println(ints())  //1
	fmt.Println(ints())  //2
	fmt.Println(ints())  //3
	fmt.Println(ints())  //4
	fmt.Println(ints())  //5

    otherints :=integers()   //別のクロージャーの定義
	fmt.Println(otherints())  //1
	fmt.Println(otherints())  //2
	fmt.Println(otherints())  //3
	fmt.Println(otherints())  //4
	fmt.Println(otherints())  //5


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

	//// int型について
	var i10 int = 1000 //環境依存(PC)でビットが指定されている
	//int8 int16 int32 int64がある
	fmt.Println(i10)
	var i11 int64 = 2000
	// fmt.Println(i10+i11)  //明示型intと推論型は計算できない
	fmt.Printf("%T\n", i10)        //int
	fmt.Printf("%T\n", i11)        //int64
	fmt.Printf("%T\n", int32(i10)) //int32
	fmt.Println(int64(i10) + i11)  //3000

	////浮動小数点数型
	var fl64 float64 = 3.4
	fmt.Println(fl64)
	fl := 3.2 //環境依存ではなく強制的にfloat64になる
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 3.2 //基本的にfloat32は使わない
	fmt.Printf("%T\n", fl32)
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	p_inf := 1.0 / zero
	fmt.Println(p_inf) //+Inf  +の無限大
	n_inf := -1.0 / zero
	fmt.Println(n_inf) //-Inf  -の無限大
	nan := zero / zero
	fmt.Println(nan) //NaN

	////string型
	var hello string = "Hello Golang"
	fmt.Println(`Golang
	Golang
	            Golang`) //バッククォーツで複数行の出力ができる
	fmt.Println("\"") //"
	fmt.Println(`"`)  //"
	//文字列はバイト型の配列になっているので
	fmt.Println(hello[0])         //72   これはバイト型で表示されてしまう
	fmt.Println("String"[0])      //83
	fmt.Println(string(hello[0])) //H

	////byte型
	byteA := []byte{72, 73}
	fmt.Println(byteA)         //[72,73]
	fmt.Println(string(byteA)) ////"HI"アスキーコード

	c := []byte("HI")
	fmt.Println(c)         //[72,73]
	fmt.Println(string(c)) //"HI"
	//アスキーコード
	//ASCII（アスキー、情報交換用米国標準コード、英: American Standard Code for Information Interchange）は、
	//現代英語や西ヨーロッパ言語で使われるラテン文字を中心とした文字コード。これはコンピュータその他の通信機器において最もよく使われているものである。

	////配列型
	//後から要素数を変更できない(C#と同様！！)
	var arr1 [3]int
	fmt.Println(arr1)        //[0 0 0]
	fmt.Printf("%T\n", arr1) //[3]int  要素数までを含めて型という扱いになる

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) //[A B ]
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	arr2[2] = "C"
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        //[C D]
	fmt.Printf("%T\n", arr4) //[2]string

	// var arr5 [4]int
	// arr5= arr1  要素数が異なるものは代入できない
	fmt.Println(len(arr1))

	//スライス型は要素数を変更可能

	////interface型 あらゆる型と互換性がある  ただし計算はできない、あくまでもあらゆる方を示すことができることに意味がある
	var x interface{}
	fmt.Println(x) //<nil>  初期値
	x = 1
	fmt.Println(x)
	x = 3.6
	fmt.Println(x)
	x = "すごい"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	////型変換
	var i20 int = 1
	fl64_ := float64(i)
	fmt.Printf("i20 = %T\n", i20)
	fmt.Printf("fl64_ = %T\n", fl64_)

	var i21 float64 = 10.5
	fmt.Println(int(i21)) //10  小数点は切り捨てて整数になる

	var s4 string = "100"
	i22, _ := strconv.Atoi(s4) //_にも値が入っているが破棄していることを記している
	fmt.Println(i22 + 11)
	fmt.Printf("i22 = %T\n", i22)
	var s5 string = "100"

	i23, err := strconv.Atoi(s5) //errを入れるとエラーハンドリングの対応になる
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i23 + 900)

	var i24 int = 400
	s6 := strconv.Itoa(i24)
	fmt.Println(s6)
	fmt.Printf("s6 = %T\n", s6)

	var h string = "Hello Golang"
	b := []byte(h)
	fmt.Println(b)

	////定数  基本はグローバルで書くことが多い
	//頭文字を大文字にするとPublicな宣言となり他のパッケージからも利用できる
	fmt.Println(Pi)
	fmt.Printf("Pi = %T\n", Pi)
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F) //1 1 1 A A A

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
