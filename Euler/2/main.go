package main

import "fmt"

func main() {
	i := 1
	sum := 0
	for {
		f := fib(i)
		if f < 4000000 && f%2 == 0 {
			sum += f
		} else if f > 4000000 {
			fmt.Println(sum)
			break
		}
		i++
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
