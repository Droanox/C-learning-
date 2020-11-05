package main

import (
	"fmt"
	"strconv"
)

func main() {
	var solution int
	for i := 100; i <= 999; i++ {
		for a := 100; a <= 999; a++ {
			if i*a > solution && isPalindrome(strconv.Itoa(i*a)) == true {
				solution = i * a
			}
		}
	}
	fmt.Println(solution)
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

//Find the largest palindrome made from the product of two 3-digit numbers.
