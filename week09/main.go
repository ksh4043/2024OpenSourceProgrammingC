package main

import (
	"fmt"

	"math/rand"
)

func main() {
	//fmt.Printf("%f 나누기 %f은(는) %f 입니다\n", 1.0, 3.0, 1.0/3.0)
	result := fmt.Sprintf("%0.2f 나누기 %0.2f은(는) %0.2f 입니다\n", 1.0, 3.0, 1.0/3.0) //Sprintf는 서식을 문자열로 리턴
	fmt.Print(result)

	i := 1
	// fmt.Printf("%T\n", i)
	// fmt.Println(reflect.TypeOf(i))
	for i <= 10 {
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		i++
	}
}
