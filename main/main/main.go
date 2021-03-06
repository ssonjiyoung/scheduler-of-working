package main

import (
	"fmt"
	"strconv"
	"time"
)

var str []int
var num = 0
var dayGoalTime, weekGoalTime, pattern, resultWeek int
var sun, mon, tus, wed, thu, fri, sat string
var selected []int

// depth, day, week, ? , defualt뺀수
func full(depth int) {
	i := 0
	fmt.Print("스탭1")
	if pattern == depth {
		fmt.Print("스탭2")
		result := 0
		var s []int
		for i = 0; i < pattern; i++ {
			fmt.Print("스탭3")
			result += selected[i]
			s = append(s, selected[i])
		}

		if result == resultWeek {
			fmt.Print("스탭4")
			sliceA := append(s, str...)
			fmt.Print(sliceA)
			num++
		}
		return
	}
	for i := 0; i <= dayGoalTime; i++ {
		fmt.Print("스탭5")
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

	selected = make([]int, pattern)

	startTime := time.Now()
	// depth
	full(0)
	fmt.Println("num = ", num)

	elapsedTime := time.Since(startTime)

	fmt.Println("실행시간 = ", elapsedTime)

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
