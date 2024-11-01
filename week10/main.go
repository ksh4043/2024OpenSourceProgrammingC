package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Number : ")
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)

	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	// 1이 소수로 나오는 bug fix
	if n <= 1 { // A Prime number is a natural number greater than 1 that has only 1 and itself as divisors
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false
				break // performance up
			}
			fmt.Printf("%d ", j) // Check j loop
			j++
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number.", n)
	}

}
