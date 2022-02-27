package main

//パッケージの宣言は必ず１つ mainの時はmainプログラムが書かれていることを指す

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println(time.Now())
}
