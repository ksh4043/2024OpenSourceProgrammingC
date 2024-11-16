package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

// git 주소에서 모듈을 가져올 때는 go get [파일주소]로 모듈을 다운로드 받고 사용
func main() {
	var gpas [3]float64
	for i := 0; i < len(gpas); i++ {
		fmt.Print("Input float number: ")
		gpas[i], _ = keyboard.GetFloat()
	}

	for _, gpa := range gpas {
		fmt.Println(gpa)
	}
}
