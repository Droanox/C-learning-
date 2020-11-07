package main

import "fmt"

func main() {

	var n int
	fmt.Print("Enter an integer from 0-1000: ")
	fmt.Scanln(&n)
	fmt.Print("Your integer in written English is:")

	if n == 0 {
		fmt.Println(" Zero")
	}

	if n > 0 && n <= 20 {
		fmt.Print(numbers[n])
	}

	if n > 20 && n < 100 {
		fmt.Print(numbers[(int(n/10))*10], numbers[n%10])
	}

	if n >= 100 && n < 1000 {
		fmt.Print(numbers[int(n/100)], " hundred")

		if n%100 != 0 {
			fmt.Print(" and", numbers[(int((n%100)/10))*10], numbers[(n%100)%10])
		}
	}

	if n >= 1000 {
		fmt.Println(" Enter a valid integer retard")
	}
}

var (
	numbers = map[int]string{
		1:  " one",
		2:  " two",
		3:  " three",
		4:  " four",
		5:  " five",
		6:  " six",
		7:  " seven",
		8:  " eight",
		9:  " nine",
		10: " ten",
		11: " eleven",
		12: " twelve",
		13: " thirteen",
		14: " fourteen",
		15: " fifteen",
		16: " sixteen",
		17: " seventeen",
		18: " eighteen",
		19: " nineteen",
		20: " twenty",
		30: " thirty",
		40: " forty",
		50: " fifty",
		60: " sixty",
		70: " seventy",
		80: " eighty",
		90: " ninety",
	}
)
