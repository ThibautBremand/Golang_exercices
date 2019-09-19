/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program, “findian.go”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	fmt.Println("Please enter a string !")

	// Use of scanner instead of fmt.Scanln because the fmt package intentionally filters out whitespaces which is not acceptable in our case.
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()

		// Not case-sensitive
		input = strings.ToLower(input)
		fmt.Println("The value you have entered:", input)

		// Check if the first character is 'i', the last one is 'n' and the string contains an 'a'
		//if (string(input[0]) == "i" && string(input[len(input) - 1]) == "n" && strings.Contains(input, "a")) {
		if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
