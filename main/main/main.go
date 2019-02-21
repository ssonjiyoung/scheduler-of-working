package main

import (
	"fmt"
)

var selected [3]int

func full(depth int) {
	input1 := 4
	input2 := 3

	i := 0
	if input2 == depth {
		for i = 0; i < input2; i++ {
			fmt.Print(selected[i])
		}
		fmt.Println()
		return
	}
	for i := 0; i < input1; i++ {
		selected[depth] = i
		full(depth + 1)
	}
}

func main() {
	full(0)
	return
}

// func RPerm(set []int, set_size int, m int, n int) {
// 	// 종료
// 	for set_size == n {
// 		fmt.Println(set, set_size)
// 		return
// 	}
// 	// 분기
// 	for i := 0; i < m; i++ {
// 		set[set_size] = i
// 		RPerm(set, set_size+1, m, n)
// 	}
// }

// func swap(a int, b int) {
// 	tmp := a
// 	a = b
// 	b = tmp
// }

// func print_arr(size int) {
// 	for i := 0; i < size; i++ {
// 		fmt.Println(arr[i])
// 	}
// 	fmt.Println()
// }

// func check_arr(size int) int{
// 	tot := 0
// 	for i := 0; i < size; i++ {
// 		if arr[i] > input2 {
// 			return 0;
// 		}
// 		tot +=
// 	}
// 	fmt.Println()
// }
