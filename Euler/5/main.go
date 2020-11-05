package main

import (
	"fmt"
	"os"
)

func main() {
	for n := 20; n >= 0; n = n + 20 {
		var sum int
		for i := 1; i <= 20; i++ {
			if n%i == 0 {
				sum += 1
			}
			if sum == 20 {
				fmt.Println(n)
				os.Exit(0)
			}
		}
		sum = 0
	}
}
