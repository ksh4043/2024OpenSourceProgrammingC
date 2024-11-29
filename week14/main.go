package main

import "fmt"

func main() {
	// 이름, 나이 입력 -> Q 입력 시 종료
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("What is your name? (exit to 'q') ")
		fmt.Scanln(&name)
		if name == "q" || name == "Q" {
			break
		}
		fmt.Print("Ur age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
