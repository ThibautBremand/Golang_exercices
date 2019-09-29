/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	 |   Food eaten	 |   Locomotion method	|   Spoken sound
---------------------------------------------------------------
cow	     |   grass	     |   walk	            |   moo
bird	 |   worms	     |   fly	            |   peep
snake	 |   mice	     |   slither	        |   hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three
fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const sep = " "
const nbArgsPerRequest = 2

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {

	// Instanciate the Animal objects
	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	empty := Animal{
		food:       "",
		locomotion: "",
		noise:      "",
	}

	// Loop to prompt the user
	for true {
		fmt.Println("> Please enter an animal name and action, separated by a space")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			request := scanner.Text()

			// Not case sensitive
			splittedReq := strings.Split(strings.ToLower(request), sep)

			// Retrieve the animal name and action from the user request
			if len(splittedReq) == nbArgsPerRequest {
				currAnimal := strings.TrimSpace(splittedReq[0])
				currAction := strings.TrimSpace(splittedReq[1])

				if currAnimal == "cow" || currAnimal == "bird" || currAnimal == "snake" {
					// Switch case on the animal name
					moveFn := GenAnimalActionFn(empty)
					switch currAnimal {
					// A generated function, depending upon the given animal, is set to moveFn
					case "cow":
						moveFn = GenAnimalActionFn(cow)
					case "bird":
						moveFn = GenAnimalActionFn(bird)
					case "snake":
						moveFn = GenAnimalActionFn(snake)
					}

					// Call the generated function with the action as parameter
					fmt.Println(moveFn(currAction))
				}
			}
		}
	}
}

/*
 * This function generates an anonymous function depending upon the given Animal as parameter.
 * The anonymous function takes an action as parameter, and returns the result of the action on the actual animal
 * Params
 * in : a (Animal), the animal that needs to have its move called
 * Return : a function set for this kind of animal, allowing us to call its action
 */
func GenAnimalActionFn(a Animal) func(move string) string {
	return func(move string) string {
		switch move {
		case "food":
			return a.Eat()
		case "locomotion":
			return a.Move()
		case "noise":
			return a.Speak()
		}
		return ""
	}
}
