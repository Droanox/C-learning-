package main

import (
	"fmt"
	"math"
)

func main() {
	variable := 0.0
	fmt.Print("Enter value to get all prime factors: ")
	fmt.Scan(&variable)
	fmt.Println("Array of prime factors:", primeFactors(variable))
}

func primeFactors(n float64) (arr []int) {
	for math.Mod(n, 2) == 0 {
		arr = append(arr, 2)
		n = n / 2
	}

	for i := 3.0; i <= math.Sqrt(n); i = i + 2 {
		if math.Mod(n, i) == 0 {
			arr = append(arr, int(i))
			n = n / i
		}
	}

	if n > 2 {
		arr = append(arr, int(n))
	}

	return arr
}
