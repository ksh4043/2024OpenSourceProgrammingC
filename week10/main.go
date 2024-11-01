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
	in := bufio.NewReader(os.Stdin) //os.Stdin은 키보드 입력 개체 반환
	fmt.Print("Input Number : ")
	i, err := in.ReadString('\n')

	if err != nil { //nil은 정상 return 값
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)  // 입력 받은 i 값의 줄바꿈(공백) 제거. 파이썬의 lstrip, rstrip 메서드와 유사
	n, err := strconv.Atoi(i) // 문자열을 다른 타입으로 convert 해줌, ParseInt의 경우 파라미터 3개 필요(변환할 문자열, 바꿀 수의 진수, 비트 크기)

	if err != nil { //nil은 정상 return 값
		log.Fatal(err)
	}

	counts := 0

	j := 2
	for j < n {
		if n%j == 0 {
			counts++
		}
		j++
	}

	if counts == 0 {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number.", n)
	}

}
