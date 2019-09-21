/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

Submit your source code for the program, “makejson.go”.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	var myMap map[string]string

	// Get the name
	fmt.Println("Please enter a name !")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()

		// Get the address
		fmt.Println("Please enter an address !")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			address = scanner.Text()

			// Init the map
			myMap = make(map[string]string)

			// Add the name & address to the map
			myMap["name"] = name
			myMap["address"] = address

			// Transform the map into jSON
			barr, err := json.Marshal(myMap)

			//  Print the jSON as string
			if err == nil {
				fmt.Println(string(barr))
			}
		}
	}
}
