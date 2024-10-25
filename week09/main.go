package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //실행시마다 time객체로 시간값을 seed값으로 설정 >> 매 실행마다 다른 값 추출 유도
	target := rand.Intn(6)       //0~5
	fmt.Println(target)
}
