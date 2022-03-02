package pai

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
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

//スライスから特定の要素を削除する関数
func Remove_str(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

//スライスから特定の要素を削除する関数 int
func Remove_int(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

//スライスから特定の要素を含んでいるかどうかを返す関数
func Contains(list interface{}, target interface{}) (bool, error) {

	switch list.(type) {
	default:
		return false, fmt.Errorf("%v is an unsupported type", reflect.TypeOf(list))
	case []int:
		revert := list.([]int)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []uint64:
		revert := list.([]uint64)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil

	case []string:
		revert := list.([]string)
		for _, r := range revert {
			if target == r {
				return true, nil
			}
		}
		return false, nil
	}

	return false, fmt.Errorf("processing failed")
}

//スライスの最大値を取得
func MaxInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[len(a)-1]
}

//スライスの最小値を取得
func MinInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}
