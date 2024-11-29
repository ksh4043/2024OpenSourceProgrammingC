package main

import "fmt"

type student struct { // 구조체를 type 키워드로 정의 -> 하나의 리터럴 타입이 됨
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student

	student1.id = 202444091
	student1.name = "김시현"
	student1.gpa = 4.5

	fmt.Println(student1.gpa)

	var student2 student

	student2.id = 202444021
	student2.name = "아이다"
	student2.gpa = 4.43

	fmt.Println(student2.id)
}
