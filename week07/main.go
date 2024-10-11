package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Input your name : ")
	//name, _ := in.ReadString('\n') _는 결과 return 값을 받지만 사용하지 않을 때
	name, err := in.ReadString('\n')
	//fmt.Println(i, err)	//err > nil:에러가 발생하지 않았을 때 return 값
	fmt.Println(name)
	fmt.Println(err)
}
