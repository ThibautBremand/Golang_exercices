/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. 
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. 
GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. 
The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. 
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

const sep = " "
const maximumNb_firstInput = 3
const maximumNb_secondInput = 1
const stopExeChar = "x"

func main() {
	var acceleration, initVelocity, initDisplacement, time float64

	b := true
	stopExe := false
	
	// Keep looping while the user has not entered "maximumNb_firstInput" correct floats
	for b {
		b = promptFirstValues(&acceleration, &initVelocity, &initDisplacement)
	}
	
	// Display only
	fmt.Println("Entered values :")
	fmt.Println("- Acceleration : " + strconv.FormatFloat(acceleration, 'f', 6, 64))
	fmt.Println("- Initial velocity : " + strconv.FormatFloat(initVelocity, 'f', 6, 64))
	fmt.Println("- Initial displacement : " + strconv.FormatFloat(initDisplacement, 'f', 6, 64))

	for !stopExe {
		b = true
		for b && !stopExe {
			b, stopExe = promptSecondValues(&time)
		}

		if !stopExe {
			// Display only
			fmt.Println("- Time : " + strconv.FormatFloat(time, 'f', 6, 64))

			// Displacement calculation
			// 1- Generate the displacement calculation function using the three parameters
			displaceFn := GenDisplaceFn(acceleration, initVelocity, initDisplacement)

			// 2- Call the generated function using the time parameter
			fmt.Println("Displacement calculated :")
			fmt.Println(displaceFn(time))
		}
	}
}

/**
 * Prompt the user to enter three values : acceleration, initial velocity, and initial displacement
 * Params
 * in :
 * - acceleration (*float64) : the pointer to the acceleration variable
 * - initVelocity (*float64) : the pointer to the initVelocity variable
 * - initDisplacement (*float64) : the pointer to the initDisplacement variable
 * Return : inputError (bool), false if no error occured, else true
 */
func promptFirstValues(acceleration, initVelocity, initDisplacement *float64) bool {
	var userInput string
	var inputError bool
	inputError = false
	mySlice := make([]float64, 0)

	// Prompt the user
	fmt.Println("Please type in values for acceleration, initial velocity, and initial displacement ! Enter all of them separated by a space")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
		floats := strings.Split(userInput, sep)

		// Check if the user entered at least "maximumNb_firstInput" values
		if len(floats) < maximumNb_firstInput {
			inputError = true
		} else {
			// Loop through all the entered floats
			for ind, elt := range floats {
				
				// Cast the current element into a float64
				f, err := strconv.ParseFloat(strings.TrimSpace(elt), 64)
				if err == nil {

					// Check if the user entered too much values
					if len(mySlice) >= (maximumNb_firstInput) {
						fmt.Println("You have entered too much values, only the first " + strconv.Itoa(maximumNb_firstInput) + " floats have been added !")
						break
					}

					// Append the floats to the slice
					mySlice = append(mySlice, f)
				} else {
					fmt.Println("The " + (strconv.Itoa(ind+1)) + "th value hasn't been added.")
				}
			}
		}

		// Check if we managed to add "maximumNb_firstInput" values
		if len(mySlice) < maximumNb_firstInput {
			inputError = true
		}

		if inputError {
			fmt.Println("Please enter " + strconv.Itoa(maximumNb_firstInput) + " values !")
		}
	} else {
		return true
	}

	if !inputError {
		*acceleration = mySlice[0]
		*initVelocity = mySlice[1]
		*initDisplacement = mySlice[2]
	}
	return inputError
}

/**
 * Prompt the user to enter one value : time
 * Params
 * in :
 * - time (*float64) : the pointer to the time variable
 * Return : 
 * - inputError (bool), false if no error occured, else true
 * - stopExe (bool), true if the user wants to stop the execution, else false
 */
func promptSecondValues(time *float64) (bool, bool) {
	var userInput string
	var inputError bool
	inputError = false

	// Prompt the user
	fmt.Println("Please type in a value for time, or enter " + stopExeChar + " to quit !")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()

		// Cast the value into a float64
		f, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
		if err == nil {
			*time = f
		} else {
			if strings.TrimSpace(strings.ToLower(userInput)) == stopExeChar {
				return true, true
			} else {
				fmt.Println("The value hasn't been added.")
				return true, false
			}
		}
	} else {
		return true, false
	}
	return inputError, false
}

/**
 * Generate a displacement calculation function from the given acceleration, initial velocity and initial displacement
 * The generated function will take a time as parameter to allow the complete displacement calculation
 * Params
 * in :
 * - a (float64) : acceleration
 * - vo (float64) : initial velocity
 * - so (float64) : initial displacement
 * Return : a function which computes displacement as a function of time, assuming the given values acceleration (a), initial velocity (vo), and initial displacement (so). (s =½ a t2 + vot + so). 
 * 			It takes takes one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.
 */
func GenDisplaceFn(a, vo, so float64) func (t float64) float64 {
	return func (t float64) float64 {
		// s =½ a t2 + vot + so
		return (a/2) * t * 2 + vo * t + so
	}
}