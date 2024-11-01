package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { // All even numbers except 2 are not prime
		isPrime = false
	} else {
		j := 3 // start value
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break // performance up
			}
			fmt.Printf("%d ", j) // Check j loop
			j = j + 2            // increment
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is not prime number.", n)
	}

}
