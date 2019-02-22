package main

import (
	"fmt"
)

var selected [5]int
var num = 0

func full(depth int) {
	// day
	input1 := 9
	// ?
	input2 := 5
	// week
	input3 := 35
	i := 0

	if input2 == depth {
		result := 0
		var s []int
		for i = 0; i < input2; i++ {
			result += selected[i]
			s = append(s, selected[i])
			// if result == input1 {
			// 	fmt.Print("s =", selected[i])
			// }
			//fmt.Print(selected[i])
		}

		if result == input3 {
			fmt.Println(s)
			num++
		}
		//fmt.Println()

		return
	}
	for i := 0; i <= input1; i++ {
		selected[depth] = i
		full(depth + 1)
	}
}

func main() {
	full(0)
	fmt.Println("num = ", num)
	return
}
