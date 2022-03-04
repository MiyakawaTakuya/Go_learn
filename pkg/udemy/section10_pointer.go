package udemy

import (
	"fmt"
	// "time"
)

//ポインタとは、値型に分類されるデータ構造(基本型、参照型、構造体）のメモリ上のアドレスと型の情報のことです。
func double(i int) {
	i = i * 2
}

func double2(i *int) {
	*i = *i * 2
}

func double3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}
func Section10_pointer() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) //0xc0000b2008 メモリ上のアドレス
	double(n)
	fmt.Println(n) //100になる
	fmt.Println(&n)

	//ポインタ型
	var p *int = &n //アスタリスクと型でポインタ型を指定している  &はアドレス演算子
	fmt.Println(p)  //0xc0000b2008
	fmt.Println(*p) //100  実体を示す
	// *p = 300
	// fmt.Println(n) //300
	// n = 200
	// fmt.Println(*p) //200

	double2(&n)    //
	fmt.Println(n) //200になる
	fmt.Println(&n)
	double2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	double3(sl)
	fmt.Println(sl) //[2 4 6]  スライスは参照型の為、値が変わる

}
