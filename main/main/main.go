package main

import (
	"fmt"
)

func main() {
	var m, n int

	m = 3
	n = 3

	//생성
	var arr []int
	for i := 0; i < m; i++ {
		arr[i] = i
	}

	fmt.Println("중복순열 출력")
	RPerm(arr, 0, m, n)
}

func RPerm(set []int, set_size int, m int, n int) {
	// 종료
	for set_size == n {
		fmt.Println(set, set_size)
		return
	}
	// 분기
	for i := 0; i < m; i++ {
		set[set_size] = i
		RPerm(set, set_size+1, m, n)
	}
}
