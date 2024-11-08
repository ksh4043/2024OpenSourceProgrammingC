package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 { // A Prime number is a natural number greater than 1 that has only 1 and itself as divisors
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 { // All even numbers except 2 are not prime
		return false
	} else {
		j := 3 // start value
		//for j <= int(math.Sqrt(float64(n))) {
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			//fmt.Printf("%d ", j) // Check j loop
			j = j + 2 // increment
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input Start Number : ")
	a, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input End Number : ")
	b, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)

	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}

}
