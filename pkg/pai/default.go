package pai

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "math"
	// "time"
)

func C() {
	fmt.Println()
}

var sc = bufio.NewScanner(os.Stdin)

//読み込み関数 一式
//文字列取得・1数値
func GetInt() (IntReturned int) {
	sc.Scan()
	IntReturned, _ = strconv.Atoi(sc.Text())
	return
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

//スペース区切りのstring配列データの取り方
func GetStrArr() []string {
	sc.Scan()
	s := sc.Text()               // 1行読み込み
	arr := strings.Split(s, " ") // 1行をスペースで分割
	return arr
}

func GetStr() (str string) {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n') //改行が現れるまでのstringを読み込む
	str = strings.TrimSpace(s)      //余分なものを取り除いてあげる 改行
	return
}
