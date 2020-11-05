package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(whatPrime(100000))
}

func whatPrime(n float64) (i float64) {
	var sum float64
	for i := 2.0; i >= 0; i++ {
		if isPrime(i) == true {
			sum += 1
			if sum == n {
				return i
			}
		}
	}

	return 0
}

func isPrime(n float64) bool {
	var sum int
	for i := 1.0; i <= math.Sqrt(n); i++ {
		if math.Mod(n, i) == 0 {
			sum += 1
		}
	}
	if sum == 1 {
		return true
	} else {
		return false
	}
}
