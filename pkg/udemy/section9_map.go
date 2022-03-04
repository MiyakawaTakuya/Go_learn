package udemy

import (
	"fmt"
)

func Section9_map() {
	////map  ハッシュ（連想配列）
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"C": 1000, "D": 2000}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string) //[]
	fmt.Println(m4)
	m4[1] = "JAPAN" //map[1:JAPAN 2:JAPON]
	m4[2] = "JAPON"
	fmt.Println(m4)
	fmt.Println(m["A"]) //100
	fmt.Println(m4[2])  //JAPON
	fmt.Println(m4[3])  //空 nilでもない  →エラーハンドリングの機能で対処する

	//取り出し
	s := m4[1]
	fmt.Println(s)
	s, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	m4[2] = "RUSIA"
	fmt.Println(m4)
	fmt.Println(len(m4))

	//for とmap
	for k, v := range m4 {
		fmt.Println(k, v)
	}
	for k := range m4 {
		fmt.Println(k)
	}
	for _, v := range m4 {
		fmt.Println(v)
	}

}
