/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var wrongInput bool
	wrongInput = true

	// Keep looping while the entered input is not a correct float
	for wrongInput {
		fmt.Println("Please enter a floating point number, it'll be displayed as an integer !")
		fmt.Scanln(&input)

		// Try to convert the entered input into a float. If it fails, the program asks the user to enter another value
		f, err := strconv.ParseFloat(input, 64)
		if err == nil {
			/** displayg the type of the b variable */
			//fmt.Printf("Type: %T \n", f)

			/** displaying the string variable into the console */
			fmt.Println("The value you have entered:", f)
			fmt.Println("The integer value:", int(f))

			// Tell the program that the user did enter a correct value, so we don't have to ask the user for another value again
			wrongInput = false
		} else {
			fmt.Println("Please enter a float !")
		}
	}
}
