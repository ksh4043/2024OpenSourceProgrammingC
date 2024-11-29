package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calc_cost(visitors []visitor) int {
	// visitors : 구조체 슬라이스
	total_cost := 0

	for _, value := range visitors {
		total_cost = total_cost + value.cost
	}

	return total_cost
}

func main() {
	var num_visitors int
	fmt.Println("How many visitors? ")
	fmt.Scanln(&num_visitors)

	vs := make([]visitor, num_visitors) // create Slice

	for i := 0; i < num_visitors; i++ {
		var age int
		fmt.Print("Input age : ")
		fmt.Scan(&age)

		if age < 12 {
			vs[i] = visitor{age: age, cost: 5000}
		} else if age >= 12 && age < 65 {
			vs[i] = visitor{age: age, cost: 10000}
		} else {
			vs[i] = visitor{age: age, cost: 7000}
		}
	}

	fmt.Printf("Total price is %d", calc_cost(vs))
}
