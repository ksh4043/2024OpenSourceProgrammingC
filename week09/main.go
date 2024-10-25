package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //실행시마다 time객체로 시간값을 seed값으로 설정 >> 매 실행마다 다른 값 추출 유도
	answer := rand.Intn(6) + 1   //1~6
	fmt.Println(answer)

	in := bufio.NewReader(os.Stdin) //os.Stdin은 키보드 입력 개체 반환
	fmt.Print("점수 입력 : ")
	i, err := in.ReadString('\n')

	if err != nil { //nil은 정상 return 값
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	guess, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답입니다.")
	} else if answer > guess {
		fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
	}
}