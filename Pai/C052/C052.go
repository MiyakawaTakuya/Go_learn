package main

//再チャレンジ C052:ゲームの画面
//math.Abs() https://pkg.go.dev/math#Abs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	// "time"
)

var H, W, dy, dx, sum int64
var line []int

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// s, _ := reader.ReadString('\n')
	// t := strings.Split(s, " ")
	// H =
	
	//VScodeはこれで通るのだけども....
	// line = GetIntArr()
	// H = int64(line[0])
	// W = int64(line[1])
	// line = GetIntArr()
	// dy = int64(math.Abs(float64(line[0])))
	// dx = int64(math.Abs(float64(line[1])))
	// sum = dy*W + dx*H - (dy * dx)
	// fmt.Println(sum)
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
