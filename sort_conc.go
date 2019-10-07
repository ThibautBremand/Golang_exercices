/*
Write a program to sort an array of integers. 
The program should partition the array into 4 parts, each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. 
Each goroutine which sorts ¼ of the array should print the subarray that it will sort. 
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main
 
import (
    "fmt"
    "math/rand"
	"strconv"
	"os"
	"bufio"
	"strings"
	"sync"
)

const sep = " "

func main() {
	var wg sync.WaitGroup
	mySlice := make([]int, 0, 0)
	mySlice = loadSlice(mySlice)

	chunkSize := (len(mySlice) + 4 - 1) / 4
	divided := chunkSlice(mySlice, chunkSize)
 
    fmt.Println("--- Unsorted --- ", mySlice)

    for index, _ := range divided {
    	wg.Add(1)
    	go sortChunk(divided, index, &wg)
    }

    wg.Wait()

    quicksort(mySlice)
    fmt.Println("--- Sorted ---", mySlice)
}

func chunkSlice(mySlice [] int, chunkSize int) [][]int {
	var divided [][]int
	for i := 0; i < len(mySlice); i += chunkSize {
	    end := i + chunkSize

	    if end > len(mySlice) {
	        end = len(mySlice)
	    }

	    divided = append(divided, mySlice[i:end])
	}

	return divided
}

func loadSlice(mySlice []int) []int {
	var userInput string

	// Prompt the user
	fmt.Println("Please type in a sequence of integers ! Enter all of them separated by a space !")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
		integers := strings.Split(userInput, sep)

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

	}

	return mySlice
}

func sortChunk(divided [][]int, index int, wg *sync.WaitGroup) {
	quicksort(divided[index])
	fmt.Println("Goroutine completed : ", divided[index])
	wg.Done()
}
  
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }
      
    left, right := 0, len(a)-1
      
    pivot := rand.Int() % len(a)
      
    a[pivot], a[right] = a[right], a[pivot]
      
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
      
    a[left], a[right] = a[right], a[left]
      
    quicksort(a[:left])
    quicksort(a[left+1:])
      
    return a
}