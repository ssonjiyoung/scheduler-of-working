package main
 
import(
   "fmt"
   "strconv"
)


func main() {
	// 일주일 목표 근무시간
   weekGoalTime := 1
   // 하루 목표 근무 시간
   dayGoalTime := 1
   // 내보낼 스케줄 패턴
   outschedulerPattern := [7]int{ 0,0,0,0,0,0,0 }
   // 사용자에게 받을 스케줄 패턴
   inschedulerPattern := [7]string{ "?","?","?","?","?","?","?" }

   var x, y, z string
   fmt.Println("----------------------------------------------------")
   fmt.Println("***********일주일(7일) 근무시간 스케줄표**************")
   fmt.Println("----------------------------------------------------")
   fmt.Println("\n\n\n")

   fmt.Println("\n----------------------------------------------------")
   fmt.Println("일주일 목표 근무시간(x)을 입력하세요( 1 <= x <= 63 ) \n = > ")
   
   x, err = fmt.Scan(&weekGoalTime)
   if err != nil {
      fmt.Println("잘못 입력하셨습니다. 처음 부터 다시 시도하세요.")
      fmt.Println("실패이유 = ", err)
      return
   } else if weekGoalTime < 1 || weekGoalTime > 63 {
      fmt.Println("잘못 입력하셨습니다. 처음 부터 다시 시도하세요.")
      fmt.Println("실패이유 = 조건에 맞지 않습니다. ( 1 <= x <= 63 )")
      return
   }

   fmt.Println("\n----------------------------------------------------")
   fmt.Println("하루 최대 근무시간(y)을 입력하세요( 1 <= x <= 9 ) \n = > ")
   y, err = fmt.Scan(&dayGoalTime)
   if err != nil {
      fmt.Println("잘못 입력하셨습니다. 처음 부터 다시 시도하세요.")
      fmt.Println("실패이유 = ", err)
   } else if dayGoalTime < 1 || dayGoalTime > 9 {
      fmt.Println("잘못 입력하셨습니다. 처음 부터 다시 시도하세요.")
      fmt.Println("실패이유 = 조건에 맞지 않습니다. ( 1 <= x <= 9 )")
      return
   } else if dayGoalTime > weekGoalTime {
	fmt.Println("잘못 입력하셨습니다. 처음 부터 다시 시도하세요.")
	fmt.Println("실패이유 = 하루 근무시간이 일주일 목표 근무시간보다 큼니다. ")
	return
   }

   fmt.Println("\n----------------------------------------------------")
   fmt.Println("스케줄 패턴을 입력하세요")
   fmt.Println("( 1 <= default 수 <= 9, 나머지 ?로 적어주세요)")
   fmt.Println("( ex. ? ? ? 8 ? ? ? ) \n = > ")
   z, err = fmt.Scan(&inschedulerPattern)
   if err != nil {
      fmt.Println("잘못 입력하셨습니다. 다시 시도하세요.")
      fmt.Println("실패이유 = ", err)
   } 

   randomNum := 0
   result := weekGoalTime
   for i, inschedulerPattern := range inschedulerPattern {
      // default 숫자라면
	  outschedulerPattern[i], err = strconv.Atoi(inschedulerPattern[i])
	  result = result - strconv.Atoi(inschedulerPattern[i])
      // 숫자가 아니라면
    	if err != nil {
			// ?가 아니라면
			if inschedulerPattern[i] != "?" {
				fmt.Println("잘못 입력하셨습니다. 다시 시도하세요.")
				fmt.Println("실패이유 = 스케줄 패턴을 숫자와 ?로 입력해주세요.")
				return
			}
			// defualt 10이지만 다시 한번 처리
			outschedulerPattern[i] = 10
			randomNum ++
			for j =0,j<len(randomNum);j++ {
				for j =0,j<=len(dayGoalTime);j++ {
					result - randomNum
					outschedulerPattern[i] = j
				}
			}
   		}
   }


   
}

// func RandomFor(int dayGoalTime, int i) {
	
// }