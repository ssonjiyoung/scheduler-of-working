package main

import (
	"fmt"
	"reflect"
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

	_, err := fmt.Scan(&weekGoalTime)
	if err != nil {
		fmt.Println("잘못 입력함. 실패이유 = ", err)
		return
	} else if weekGoalTime < 1 || weekGoalTime > 63 {
		fmt.Println("잘못 입력함. 조건 :  1 <= x <= 63 ")
		return
	}

	fmt.Println("하루 최대 근무시간(y) 입력( 1 <= x <= 9 ) ↓ ")
	_, err = fmt.Scan(&dayGoalTime)
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
	_, err = fmt.Scan(&sun, &mon, &tus, &wed, &thu, &fri, &sat)
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
	daytype := reflect.TypeOf(day)
	if day == "?" {
		pattern++
	} else if daytype.Name() == "int" {
		str = append(str, num)
		// default 숫자라면
		resultWeek -= num
		fmt.Println(resultWeek)
	} else {
		fmt.Println("잘못 입력함. 조건 : 스케줄 패턴을 숫자와 ?로 입력")
		return
	}
}
