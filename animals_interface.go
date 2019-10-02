/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. 
With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. 
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal	 |   Food eaten	 |   Locomotion method	|   Spoken sound
---------------------------------------------------------------
cow	     |   grass	     |   walk	            |   moo
bird	 |   worms	     |   fly	            |   peep
snake	 |   mice	     |   slither	        |   hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. 
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. 
The third string is the type of the new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. 
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. 
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. 
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. 
Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const sep = " "
const nbArgsPerRequest = 3
const commandQuery = "query"
const commandCreateAnimal = "newanimal"

var cowData = [...]string {"grass", "walk", "moo"}
var birdData = [...]string {"worms", "fly", "peep"}
var snakeData = [...]string {"mice", "slither", "hsss"}

type Animal interface {
    Eat()
    Move()
    Speak()
}

//******* COW **********
type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

//******* BIRD **********
type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

//******* SNAKE **********
type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	myMap := make(map[string]Animal)

	// Loop to prompt the user
	for true {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			request := scanner.Text()
			// Not case sensitive
			splittedReq := strings.Split(strings.ToLower(request), sep)
			if len(splittedReq) == nbArgsPerRequest {
				query := strings.TrimSpace(splittedReq[0])
				arg1 := strings.TrimSpace(splittedReq[1])
				arg2 := strings.TrimSpace(splittedReq[2])

				switch query {
				case commandQuery:
					// The user wants to query the map of Animals
					handleQuery(myMap, arg1, arg2)
				case commandCreateAnimal:
					// The user wants to create a new Animal in the map
					myMap = handleCreateAnimal(myMap, arg1, arg2)
				default:
					fmt.Println("Unknown command !")
				}
			}
		}
	}
}

/**
 * This function calls the printing function
 * Params
 * - in : 
 * myMap, the map of Animals
 * animalName (string), the animal name entered by the user
 * animalInfo (string), the animal information the user wants to see
 */
func handleQuery(myMap map[string]Animal, animalName string, animalInfo string) {
	PrintAnimalInfo(myMap[animalName], animalInfo)
}

/**
 * This function calls the function that creates the Animals, and add the created Animal into the map of Animals
 * Params
 * - in :
 * myMap, the map of Animals
 * animalName (string), the animal name entered by the user
 * animalType (string), the type of the animal
 */
func handleCreateAnimal(myMap map[string]Animal, animalName string, animalType string) map[string]Animal {
	a := CreateAnimal(animalType)
	if a != nil {
		myMap[animalName] = a
	} else {
		fmt.Println("Not added !")
	}
	return myMap
}

/**
 * This function displays the wanted information about a given Animal
 * Params
 * - in :
 * a (Animal), the Animal
 * animalInfo (string), the information we want to display about the animal
 */
func PrintAnimalInfo(a Animal, animalInfo string) {
	if a != nil {
		switch animalInfo {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			fmt.Println("Wrong command !")
		}
	} else {
		fmt.Println("Wrong name !")
	}
}

/**
 * This function creates a new Animal using a given specie
 * Params
 * - in :
 * animalType (string), the animal specie (cow, bird or snake)
 * Return :
 * - a new Animal, or nil if a wrong specie is given
 */
func CreateAnimal(animalType string) Animal {
	switch animalType {
	case "cow":
		return Cow{cowData[0], cowData[1], cowData[2]}
	case "bird":
		return Bird{birdData[0], birdData[1], birdData[2]}
	case "snake":
		return Snake{snakeData[0], snakeData[1], snakeData[2]}
	}
	return nil
}