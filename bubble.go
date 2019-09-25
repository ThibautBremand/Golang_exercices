/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. 
The program should print the integers out on one line, in sorted order, from least to greatest. 
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. 
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. 
You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers 
and an index value i which indicates a position in the slice. 
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
)

const sep = ","
const maximumNb = 10

func main() {
	var userInput string
	mySlice := make([]int, 0, maximumNb)

	// Prompt the user
	fmt.Println("Please type in a sequence of up to " + strconv.Itoa(maximumNb) + " integers ! Enter all of them separated by a '" + string(sep) + "'")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
		integers := strings.Split(userInput, sep)

		// Check if the user entered too much integers
		if len(integers) > maximumNb {
			fmt.Println("You have entered too much integers !")
		} else {
			// Loop through all the entered integers
			for ind, elt := range integers {
				
				// Cast the current element into an integer
				f, err := strconv.Atoi(strings.TrimSpace(elt))
				if err == nil {
					// Append the integer to the slice
					mySlice = append(mySlice, f)
				} else {
					fmt.Println("The " + (strconv.Itoa(ind+1)) + "th integer hasn't been added.")
				}
			}
			fmt.Println("")
			fmt.Println("The unsorted slice :")
			fmt.Println(mySlice)
			fmt.Println("")

			BubbleSort(mySlice)
			fmt.Println("The sorted slice :")
			fmt.Println(mySlice)
		}
	}
}

/**
 * Sort a slice of ints using a BubbleSort algorithm
 * Params :
 * - in : integers, the slice of ints to be sorted
 */
func BubbleSort(integers []int) {
    n := len(integers)
    sorted := false

    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if integers[i] > integers[i+1] {
                Swap(integers, i)
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}

/**
 * Swap two elements of a slice
 * Params :
 * - in : 
 * slice, the slice which needs to swap two elements
 * i, the position of the element in the slice which needs to be swapped with its following element
 */
func Swap(slice []int, i int) {
	slice[i+1], slice[i] = slice[i], slice[i+1]
}