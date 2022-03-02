package pai

//再チャレンジ C052:ゲームの画面
// 受験結果 受験言語： Go 獲得ランク： Cランク スコア： 100点  220302
//math.Abs() https://pkg.go.dev/math#Abs

import (
	"fmt"
	"math"
	// "bufio"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)

var H, W, dy, dx, sum int64
var n int

func C052() {
	line := GetIntArr()
	H = int64(line[0])
	W = int64(line[1])
	line = GetIntArr()
	dy = int64(math.Abs(float64(line[0])))
	dx = int64(math.Abs(float64(line[1])))
	sum = dy*W + dx*H - (dy * dx)
	fmt.Println(sum)
}
