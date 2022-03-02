package main

//再チャレンジ C052:ゲームの画面
// 受験結果 受験言語： Go 獲得ランク： Cランク スコア： 100点  220302
//math.Abs() https://pkg.go.dev/math#Abs

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "time"
)

var sc = bufio.NewScanner(os.Stdin)
var H, W, dy, dx, sum int64
var n int

func main() {
	line := GetIntArr()
	H = int64(line[0])
	W = int64(line[1])
	line = GetIntArr()
	dy = int64(math.Abs(float64(line[0])))
	dx = int64(math.Abs(float64(line[1])))
	sum = dy*W + dx*H - (dy * dx)
	fmt.Println(sum)
}

//スペース区切りのint配列データの取り方
func GetIntArr() []int {
	sc.Scan()
	s := sc.Text()              // 1行読み込み
	sl := strings.Split(s, " ") // 1行をスペースで分割
	arr := []int{}
	for i := range sl {
		v, _ := strconv.Atoi(sl[i]) // 1要素ずつ整数に変換
		arr = append(arr, v)        // スライスに追加
	}
	return arr
}
