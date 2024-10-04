package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	//mySchoolAccount := 5
	//fmt.Println(mySchoolAccount)

	fmt.Println(f, b, s, i)
	fmt.Printf("%f %t %s %d\n", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print("\n\n", f <= float64(i), "\n") // comparison (true/false)
}
