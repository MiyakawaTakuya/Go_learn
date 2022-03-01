package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "time"
)

var n int
var line []int

func main() {
	n = GetInt()
	line = GetIntArr()

	fmt.Println(n)
}

//読み込み関数
//文字列取得・1数値
func GetInt() (IntReturned int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	IntReturned, _ = strconv.Atoi(scanner.Text())
	return
}

//文字列取得・1単語stiring
func GetStr() (StrReturned string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	StrReturned = scanner.Text()
	return
}

//数字列1行取得・1整数以上対応
func GetIntArr() (intReturned []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	intReturned, _ = SplitAndConvertToInt(str, " ")
	return
}

// デリミタで分割して整数値スライスを取得
func SplitAndConvertToInt(stringTargeted, delim string) (intSlices []int, err error) {
	// 分割
	stringSplited := SplitWithoutEmpty(stringTargeted, delim)

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err = strconv.Atoi(stringSplited[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}

// 空白や空文字だけの値を除去したSplit()
func SplitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}
	return
}
