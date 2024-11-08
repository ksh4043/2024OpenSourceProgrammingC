package main

import "main.go/greeting"

func main() { //greeting package는 main이 아니기 때문에 main 파일과 함수에서 호출해서 사용
	//greeting.Hello("Inha")
	//greeting.Hi("Harvard")
	greeting.EnglishGreetings("Inha") //hello와 hi를 greeting 패키지 내부에서 호출하는 EnglishGreetings를 호출해서 사용
}
