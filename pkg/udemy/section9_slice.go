package udemy

import (
	"fmt"
	// "time"
)

func Section9_slice() {
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

	////スライスをコピーするとき
	sl10 := []int{1, 2, 3, 4, 5}
	sl11 := make([]int, 5, 10)
	fmt.Println(sl11) //[0 0 0 0 0]
	n := copy(sl11, sl10)
	fmt.Println(n, sl11)  //5 [1 2 3 4 5]
	fmt.Printf("%T\n", n) //int 変数nはコピーに成功した数が返ってきている

	//for スライス
	sl12 := []string{"A", "B", "C"}
	fmt.Println(sl12)
	for i := range sl12 {
		fmt.Println(i, sl12[i])
	}
	//古典的な書き方
	for i := 0; i < len(sl12); i++ {
		fmt.Println(sl12[i])
	}

	//スライス 可変長引数
	fmt.Println(sum(1, 3, 5))
	fmt.Println(sum(1, 3, 5, 123, 543, 656564))
	sl13 := []int{1, 2, 4, 6, 8}
	fmt.Println(sum(sl13...))

}

//引数に渡された複数の整数の合計値を返す関数
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
