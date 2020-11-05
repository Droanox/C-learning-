package main

import "fmt"

func main() {
	Answer := Modular()
	fmt.Println(Answer)
}

func Modular() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}

	}

	return sum
}
