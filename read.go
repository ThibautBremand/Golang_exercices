/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found
in each struct.

Submit your source code for the program, “read.go”.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const sep = ","
const sliceInitSize = 10

type name struct {
	fname string
	lname string
}

func main() {
	var filename string

	// Create an empty slice of strucs "name"
	mySlice := make([]name, 0, sliceInitSize)

	// Get the text filename
	fmt.Println("Please enter a text filename !")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		filename = scanner.Text()

		// Open the text filename
		file, err := os.Open("./" + filename)
		if err != nil {
			log.Fatal(err)
		}
		// The file will be closed at the end
		defer file.Close()

		// Loop through each row of the file
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			row := scanner.Text()
			splittedRow := strings.Split(row, sep)

			// Add the values to the slice
			mySlice = append(mySlice, name{splittedRow[0], splittedRow[1]})
		}
	}

	// Iterate over the slice's values and print them out
	for _, itemCopy := range mySlice {
		fmt.Println("fname: " + itemCopy.fname + " - lname: " + itemCopy.lname)
	}
}
