package udemy

import (
	"fmt"
	"time"
)

//channel
//複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造

func Section9_channel() {
	//宣言・操作
	var ch1 chan int
	ch1 = make(chan int)
	// var ch2 <-chan int //明示的な受信専用
	// var ch3 chan<- int //明示的な送信専用
	ch2 := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5) //容量が5のチャネル
	fmt.Println(cap(ch3))

	ch3 <- 1              //ch3に1を送信
	fmt.Println(len(ch3)) //1
	ch3 <- 2
	ch3 <- 4
	fmt.Println(len(ch3)) //3

	i := <-ch3
	fmt.Println(i)               //1
	fmt.Println("len", len(ch3)) //2
	i2 := <-ch3
	fmt.Println(i2)              //2
	fmt.Println("len", len(ch3)) //1
	// i3 := <-ch3
	fmt.Println(<-ch3)           //4
	fmt.Println("len", len(ch3)) //0

	//キュー 先入先出 データの順番を担保するデータ構造
	ch3 <- 10
	fmt.Println("len", len(ch3)) //1
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3)) //0
	ch3 <- 20
	ch3 <- 30
	ch3 <- 40
	ch3 <- 50 //現在のバッファサイズは5つ
	ch3 <- 60 //fatal error: all goroutines are asleep - deadlock!

	//チャネルとゴルーチン
	//複数のゴルーチン間でチャネルを触る
	// ch1_ := make(chan int)
	// fmt.Println(<-ch1_)
	// ch2_ := make(chan int)
	// go reciever(ch1_) //レシーバーを起動させる
	// go reciever(ch2_)

	// i_ := 0
	// for i_ < 100 {
	// 	ch1_ <- i_
	// 	ch2_ <- i_ * 2
	// 	time.Sleep(50 * time.Millisecond)
	// 	i_++
	// }

	ch10 := make(chan int, 2)
	// close(ch10)
	ch10 <- 1 //closeされた状態だと送信できずエラー
	close(ch10)
	i10, ok := <-ch10    //closeされていても受信はできる
	fmt.Println(i10, ok) //1 true
	i12, ok := <-ch10    //closeされていても受信はできる
	fmt.Println(i12, ok) //0 false(中身ないよー)

	ch20 := make(chan int, 2)
	go reciever2("1.goroutin", ch20)
	go reciever2("2.goroutin", ch20)
	go reciever2("3.goroutin", ch20)

	i20 := 0
	for i20 < 100 {
		ch20 <- i20
		i20++
	}
	close(ch20)
	time.Sleep(3 * time.Second) //ゴルーチン処理を待つ簡易的な処理

	ch30 := make(chan int, 3)
	ch30 <- 1
	ch30 <- 2
	ch30 <- 3
	close(ch30) //closeさせておかないと余分にとろうとしてデットロックかかってしまう
	for i30 := range ch30 {
		fmt.Println(i30)
	}

	//2つのチャネルの処理を分岐させる
	ch40 := make(chan int, 2)
	ch41 := make(chan string, 2)
	// ch41 <- "A"
	// ch40 <- 100
	// ch41 <- "Hello"
	// ch40 <- 100000

	// v1 := <-ch40
	// v2 := <-ch41
	// fmt.Println(v1)
	// fmt.Println(v2)
	//チャネルに対する分岐処理を行う。必ずselect内でチャネルを指定する
	select {
	case v1 := <-ch40:
		fmt.Println(v1 + 1000)
	case v2 := <-ch41:
		fmt.Println(v2 + "!!")
	default:
		fmt.Println("どちらでもない")
	}

	ch43 := make(chan int)
	ch44 := make(chan int)
	ch45 := make(chan int)

	//reciever 今回は無名関数で
	go func() {
		for {
			i := <-ch43
			ch44 <- i * 2
		}
	}()
	go func() {
		for {
			i2 := <-ch44
			ch45 <- i2 - 1
		}
	}()

	n := 0
	//名前付きのforにするとdefaultでbreakさせることができる
L:
	for {
		select {
		case ch43 <- n:
			n++
		case i43 := <-ch45:
			fmt.Println("recieved", i43)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100{
		// 	break
		// }
	}
}

func reciever(c chan int) {
	for {
		i_ := <-c //チャネルにデータが入ったら出力
		fmt.Println(i_)
	}
}

func reciever2(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok { //チャネルにデータが入ってfalseだったらforから抜け出して出力しない
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "END")
}
