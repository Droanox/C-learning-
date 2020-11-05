package main

import "fmt"

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
