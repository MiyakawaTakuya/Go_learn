package main

import (
	"fmt"
	// "time"
)

func main() {
	////スライス 配列に似ているもの
	var sl []int //要素数を指定しない点が配列と違う
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"100", "あいう"}
	fmt.Println(sl3)

	sl4 := make([]int, 5) //メイク関数を使って生成する
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	fmt.Println(sl5[4])
	fmt.Println(sl5[2:4]) //[2,3]]2から4の手前まで
	fmt.Println(sl5[:2])  //[1,2]
	fmt.Println(sl5[2:])  //[3,4,5]
	fmt.Println(sl5[:])   //[1,2,3,4,5]

	////append make len cap
	sl_ := []int{100, 200}
	fmt.Println(sl_)

	sl_ = append(sl_, 300)
	fmt.Println(sl_)

	sl_ = append(sl_, 400, 500, 600)
	fmt.Println(sl_)

	sl2_ := make([]int, 5)
	fmt.Println(sl2_)      //[0 0 0 0 0]
	fmt.Println(len(sl2_)) //5
	fmt.Println(cap(sl2_)) //5

	sl3_ := make([]int, 5, 10) //要素数  , キャパシティ
	fmt.Println(sl3_)          //[0 0 0 0 0]
	fmt.Println(len(sl3_))     //5
	fmt.Println(cap(sl3_))     //10
	sl3_ = append(sl3_, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl3_)      //[0 0 0 0 0 1 2 3 4 5 6 7]
	fmt.Println(len(sl3_)) //12
	fmt.Println(cap(sl3_)) //20  容量が足りなくなると自動的に倍になる仕様になっている

}
