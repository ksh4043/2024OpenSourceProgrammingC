package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
) //C언에서 #include <stdio.h> 표준 입출력 함수와 유사

func main() {
	//var i int = 55
	//var f float32 = 3.99

	//var i int
	//i = 55

	i := 55
	//i := "55"
	f := 3.99

	fmt.Printf("i는 %d\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)
	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("kim si hyeon"))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
}
