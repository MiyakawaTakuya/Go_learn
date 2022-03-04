package udemy

import (
	"fmt"
	"os"
	"strconv"
	"time"
	// "strconv"
)

func Section8() { //mainが動く
	//// fmt.Println("Hello new file")
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("I don't Know")
	}

	//簡易分付きif文  条件式の前に簡易文を付けられる
	//
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x) //2 内部の変数が優先される
	}
	fmt.Println(x) //0

	////エラーハンドリング
	var s string = "100"
	i, _ := strconv.Atoi(s)
	fmt.Printf("i = %T\n", i)

	var s2 string = "百"
	i2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i2)

	//for 繰り返し処理
	fmt.Println("for文")
	// i := 0
	// for{
	// 	i++
	// 	if i==100{
	// 		fmt.Println("Now 100")
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	for i := 0; i < 14; i++ {
		if i%3 == 0 {
			continue //以下の処理を飛ばしてforの条件に戻る
		}
		fmt.Println(i)
	}

	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	// index value
	// 0 1
	// 1 2
	// 2 3
	for i, _ := range arr {
		fmt.Println(i)
	}
	// index
	// 0
	// 1
	// 2

	for _, v := range arr {
		fmt.Println(v)
	}
	// value
	// 1
	// 2
	// 3

	//map Key Valueで格納する
	m := map[string]int{"apple": 100, "banana": 200, "cinnamon": 400}

	for k, v := range m {
		fmt.Println(k, v)
	}

	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}

	//変数を局所的に指定してあげる方が望ましい
	switch n := 5; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't Know")
	}

	//列挙型ではなく演算子を使ったcaseの書き方
	//列挙型と演算子の混在はエラーになってしまい使えないので注意
	switch {
	case n > 0 && n < 10:
		fmt.Println("10未満の正の値")
	case n > 5 && n < 20:
		fmt.Println("5より大きく20未満")
	default:
		fmt.Println("I don't Know")
	}

	////型アサーション 動的に型をチェック
	anything("あいうえお")
	anything(99)

	var x2 interface{} = 6
	i3 := x2.(int) //復元したい型を.()で表記
	i3, isInt := x2.(int)
	fmt.Println(i3, isInt)
	fmt.Println(i3 * 100)

	f, isFloate64 := x2.(float64)
	fmt.Println(f, isFloate64)

	//if を使った型アサーション
	if x2 == nil {
		fmt.Println("None")
	} else if i3, isInt := x2.(int); isInt { // isInt にtrue or falseが帰ってくるので
		fmt.Println(i3, "x2 is Int")
	} else {
		fmt.Println("i don't know")
	}

	//switchを使った型アサーション こちらの方が書きやすくわかりやすいので推奨されている
	switch x2.(type) {
	case int:
		fmt.Println("x2はintです")
	case string:
		fmt.Println("x2はstringです")
	default:
		fmt.Println("I don't Know")
	}
	switch v2 := x2.(type) {
	case int:
		fmt.Println(v2, "v2はintです")
	case string:
		fmt.Println(v2, "v2はstringです")
	default:
		fmt.Println("I don't Know")
	}

	anything("いろはにほへと")
	anything(1)

	////ラベル付きfor
Loop:
	for {
		for {
			for {
				fmt.Println("スターティン♪")
				break Loop //continue Loopでも同様に利用することができる
			}
		}
	}
	fmt.Println("The End")
	//defer

	TestDefer()
	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //最後にファイルを閉じる 解放したものを閉じる処理を忘れにくい
	file.Write([]byte("Hello Golang"))

	// defer func(){  //panic()とrecover()はセットで利用する
	// 	if x3 := recover(); x3 != nil{
	// 		fmt.Println("復帰したよ")
	// 		fmt.Println(x3) // runtime error
	// 	}
	// }()
	// panic("runtime error")   //強制的にここでruntime errorを起こさせる
	// fmt.Println("Start")

	////go goroutin  並行処理が簡単に実装できる
	go sub()
	for {
		fmt.Println("Main loop")
		time.Sleep(150 * time.Millisecond)
	}

	////init

}

////mainではない関数達  同じパッケージ内で他のファイルでも最初に呼び出される
// func init() {
// 	fmt.Println("initは特別でmain関数より先に呼ばれるようになってるんよ、覚えておいて")
// }
// func init() {
// 	fmt.Println("initは複数呼ぶこともできるけど、まとめて定義するのが通常だよ")
// }

func sub() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func TestDefer() {
	defer fmt.Println("END") //関数が終了したときの処理を登録できる
	fmt.Println("START")
	fmt.Println("MIDDLE")
	//START
	// MIDDLE
	// END
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	//後から登録されたものが順に呼び出される
	// 3
	// 2
	// 1
}

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "でござる!!!!")
	case int:
		fmt.Println(v + 9999)
	}
}
