package main

import (
	"fmt"
)

func main() {
	var emptySlice []bool
	// emptySlice = make([]bool, 5)
	fmt.Printf("%#v\n", emptySlice) // Slice Zero Value >> nil 값을 할당함

	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	// gpa_slice := gpa[1:4]
	gpa[2] = 2.71                            // slice literal을 수정하면 참조하는 원본(Array literal)이 수정됨
	gpa_slice = append(gpa_slice, 4.3, 5.55) // 참조하던 Array의 크기를 벗어나면 참조하지 않게됨. 서로 다른 메모리 주소 사용
	fmt.Println(len(gpa_slice), gpa_slice, gpa)

}
