package main

import (
	"fmt"
	"strconv"
)

var str []int
var num = 0
var dayGoalTime, weekGoalTime, pattern, resultWeek int
var sun, mon, tus, wed, thu, fri, sat string

// depth, day, week, ? , defualt뺀수
func full(depth int) {
	// day
	//dayGoalTime := 9
	// week
	//weekGoalTime := 35
	// ?
	//pattern := 5
	//var selected [5]int
	fmt.Println("start")
	selected := make([]int, pattern)
	i := 0

	if pattern == depth {
		result := 0
		var s []int
		for i = 0; i < pattern; i++ {
			result += selected[i]
			s = append(s, selected[i])
		}

		if result == resultWeek {
			fmt.Println(s, str)
			num++
		}
		//fmt.Println()
		return
	}
	for i := 0; i <= dayGoalTime; i++ {
		selected[depth] = i
		full(depth + 1)
	}
}

func main() {
	fmt.Println("*일주일(7일) 근무시간 스케줄표*")
	fmt.Println()
	fmt.Println("일주일 목표 근무시간(x) 입력( 1 <= x <= 63 ) ↓ ")

	_, err := fmt.Scanln(&weekGoalTime)
	if err != nil {
		fmt.Println("잘못 입력함. 실패이유 = ", err)
		return
	} else if weekGoalTime < 1 || weekGoalTime > 63 {
		fmt.Println("잘못 입력함. 조건 :  1 <= x <= 63 ")
		return
	}

	fmt.Println("하루 최대 근무시간(y) 입력( 1 <= x <= 9 ) ↓ ")
	_, err = fmt.Scanln(&dayGoalTime)
	if err != nil {
		fmt.Println("잘못 입력함. 실패이유 = ", err)
		return
	} else if dayGoalTime < 1 || dayGoalTime > 9 {
		fmt.Println("잘못 입력함.  조건 :  1 <= x <= 9 ")
		return
	} else if dayGoalTime > weekGoalTime {
		fmt.Println("잘못 입력함. 조건 : 하루 근무시간 < 일주일 목표 근무시간 ")
		return
	}

	fmt.Println("스케줄 패턴 입력")
	fmt.Println("( 1 <= default 수 <= 9, 나머지 ?로)")
	fmt.Println("( ex. ? ? ? 8 ? ? ? ) ↓ ")
	_, err = fmt.Scanln(&sun, &mon, &tus, &wed, &thu, &fri, &sat)
	if err != nil {
		fmt.Println("잘못 입력함. 실패이유 = ", err)
	}

	resultWeek = weekGoalTime
	PatternCheck(sun)
	PatternCheck(mon)
	PatternCheck(tus)
	PatternCheck(wed)
	PatternCheck(thu)
	PatternCheck(fri)
	PatternCheck(sat)

	// depth
	full(0)
	fmt.Println("num = ", num)
	return
}

func PatternCheck(day string) {

	if day == "?" {
		pattern++
	} else {
		num, err := strconv.Atoi(day)
		if err != nil {
			fmt.Println("잘못 입력함. 조건 : 스케줄 패턴을 숫자와 ?로 입력")
			return
		}
		str = append(str, num)
		// default 숫자라면
		resultWeek -= num
		fmt.Println(resultWeek)
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
