/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/

/*
RACE CONDITION :
Because we pass the argument x by reference, it's going to be shared between both Goroutines : this is the cause of a race condition.
In a deterministic execution, we would expect to see the below result :

	goroutine 1 : 1
	goroutine 2 : 2
	done

Or this one, depending upon which goroutine finishes executing first :

	goroutine 2 : 2
	goroutine 1 : 1
	done

Because of race condition, sometimes both goroutines have already incremented the shared x variable before any goroutine could even print their result, which gives us this result :

	goroutine 1 : 2
	goroutine 2 : 2
	done

Or this one :

	goroutine 2 : 2
	goroutine 1 : 2
	done

*/

package main

import (
    "fmt"
    "time"
)

func f(from string, x *int) {
	*x = *x + 1
    fmt.Println(from, ":", *x)
}

func main() {
	x := 0

    go f("goroutine 1", &x)
	go f("goroutine 2", &x)

    time.Sleep(time.Second)
    fmt.Println("done")
}