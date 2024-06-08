package main

import (
	"fmt"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	// ADD YOUR CODE BELOW
	fmt.Printf("Air is good on %s\n", planet)
	fmt.Println("It's " + strconv.FormatBool(isTrue))
	fmt.Printf("It is %.1f degrees", temp)
}
