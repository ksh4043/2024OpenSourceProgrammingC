package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
func test(strs ...string) { // ... 키워드로 가변 매개 변수(Slice literal)
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	// fmt.Println(os.Args[1:], len(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices[1])

	for _, slice := range slices {
		fmt.Println(slice)
	}
	slices = append(slices, "forever", "!")
	fmt.Println(slices, len(slices))
	test("abc")
	test("abc", "123")
	test()
	test("abc", "123", "456")
}
