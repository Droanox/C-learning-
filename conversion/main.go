package main

import (
	"fmt"
)

var tempValue float32
var fromTemp string
var toTemp string

func main() {

	newTemp := tempCalc(tempValue, fromTemp, toTemp)

	fmt.Println("Your new temperature in the real format is: ", newTemp)

}

func tempCalc(tempValue float32, fromTemp string, toTemp string) (newTemp float32) {

	fmt.Print("Enter the temperature you want to convert: ")
	fmt.Scanln(&tempValue)

	fmt.Print("Now enter what degree you want to convert from (Kelvin, Fahrenheit, Celsius): ")
	fmt.Scanln(&fromTemp)

	fmt.Print("Enter what degree you want to convert to: ")
	fmt.Scanln(&toTemp)

	if fromTemp == "Celsius" {
		switch {
		case (toTemp == "Kelvin"):
			newTemp = tempValue + 273
		case (toTemp == "Fahrenheit"):
			newTemp = ((((tempValue - 32) * 5) / 9) + 32)
		default:
			fmt.Println("Enter a valid unit of conversion")
		}
	}
	if fromTemp == "Kelvin" {
		switch {
		case (toTemp == "Celsius"):
			newTemp = tempValue - 273
		case (toTemp == "Fahrenheit"):
			newTemp = ((((tempValue - 273) * 9) / 5) + 32)
		default:
			fmt.Println("Enter a valid unit of conversion")
		}
	}
	if fromTemp == "Fahrenheit" {
		switch {
		case (toTemp == "Kelvin"):
			newTemp = ((((tempValue - 32) * 5) / 9) + 273.15)
		case (toTemp == "Celsius"):
			newTemp = (((tempValue - 32) * 5) / 9)
		default:
			fmt.Println("Enter a valid unit of conversion")

		}
	}

	return newTemp
}
