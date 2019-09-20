/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.

The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

Submit your source code for the program, “slice.go”.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const stopChar = "x"
const sliceInitSize = 3

func main() {
	var input string

	// Create an empty integer slice of size (length) 3
	mySlice := make([]int, 0, sliceInitSize)

	// Keep looping until the user enters 'X' or 'x'
	for true {
		fmt.Println("Please enter an integer, or X to quit")
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Try to convert the entered input into a int. If it fails, the program asks the user to enter another value
		f, err := strconv.Atoi(input)
		if err == nil {
			// Add the int to the slice
			fmt.Printf("The value you have entered: %d has been added to the slice \n", f)
			mySlice = append(mySlice, f)

			// Sort the slice and print the sorted slice
			sort.Ints(mySlice)
			fmt.Println(mySlice)

			// Reset the input in order to not add the input again in the slice if the user keeps pushing the 'Enter' button
			input = ""
		} else {
			// Leave the loop if the user enters X or x
			if input == stopChar {
				break
			}
		}
	}
}
