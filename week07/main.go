package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var int int = 99 //Go언어는 예약어를 변수명으로 사용할 수 있기는 하다.
	//var a int = 1    //앞에서 int를 변수명으로 사용해서 타입지정자로 사용할 수 없다.
	a := 1
	//var fmt string = "p" 이런 현상을 Shadowing 이라고 함
	fmt.Println(int + a)

	in := bufio.NewReader(os.Stdin) //os.Stdin은 키보드 입력 개체 반환
	fmt.Print("Input your i : ")
	i, err := in.ReadString('\n')

	if err != nil { //nil은 정상 return 값
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 입력 받은 i 값의 줄바꿈(공백) 제거. 파이썬의 lstrip, rstrip 메서드와 유사
	score, _ := strconv.ParseInt(i, 16, 32) // 문자열을 다른 타입으로 convert 해줌, ParseInt의 경우 파라미터 3개 필요(변환할 문자열, 바꿀 수의 진수, 비트 크기)
	if score >= 90 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", score)
	}
}
