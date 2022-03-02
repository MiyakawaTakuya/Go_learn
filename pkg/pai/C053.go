package pai

//受験結果 受験言語： Go 獲得ランク： Cランク スコア： 80点  220302
//標準関数が不足多すぎる....そしてなぜか２問間違い

//golangにはスライス削除関数というものがありません。

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	// "math"
	// "time"
)

var x10 int = 1

func C053() {

	N := GetInt()
	N = N * 0 //使わないが為にわざわざ書かないといけないのか....？
	cards := GetStrArr()
	// fmt.Println(N)
	//x10を含んでいるかチェック
	newCards := check10(cards)
	//intのnewCards_intの配列を作成
	// newCards_int := make([]int, N)
	var newCards_int []int
	for i := range newCards {
		v, _ := strconv.Atoi(newCards[i])      // 1要素ずつ整数に変換
		newCards_int = append(newCards_int, v) // スライスに追加
	}
	// fmt.Println(newCards_int)  //[0 10 1 5 3]
	//0を含んでいるかチェック
	newCards_ans := check0(newCards_int)
	// fmt.Println(newCards_ans )
	//intCardsの合計値にx10をかけて出力する
	sum := 0
	for _, x := range newCards_ans {
		sum += x
	}
	sum *= x10
	fmt.Println(sum)

}

func check10(car []string) (cards []string) {
	a, _ := contains(car, "x10")

	if a {
		cards = remove(car, "x10")
		x10 = 10
	} else {
		cards = car
	}
	return
}

func check0(car []int) (cards []int) {
	a, _ := contains(car, 0) //0を含んでいる場合は変換
	// fmt.Println(a) //true
	if a {
		max := maxInt(car)
		// fmt.Println(max)  //10
		cards = remove_int(car, max)
	} else {
		cards = car
	}
	return
}

//スライスから特定の要素を削除する関数
func remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

//スライスから特定の要素を削除する関数 int
func remove_int(ints []int, search int) []int {
	result := []int{}
	for _, v := range ints {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

//スライスから特定の要素を含んでいるかどうかを返す関数
func contains(list interface{}, target interface{}) (bool, error) {

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
func maxInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[len(a)-1]
}

//スライスの最小値を取得
func minInt(a []int) int {
	sort.Sort(sort.IntSlice(a))
	return a[0]
}
