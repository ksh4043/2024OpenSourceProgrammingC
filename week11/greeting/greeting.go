package greeting

import "fmt"

func hi(name string) { //함수를 소문자로 작성 시 외부에서 접근이 불가능
	fmt.Printf("Hi %s\n", name)
}

func hello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func EnglishGreetings(name string) { //단 패키지 내부에서 소문자로 작성한 함수를 호출하면 외부에서 사용할 수 있음
	hello(name)
	hi(name)
}
