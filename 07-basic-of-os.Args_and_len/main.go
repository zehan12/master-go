package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args prints the command line arguments as a slice.
	// The 0th index always contains the path of the executable,
	// and the remaining elements are the arguments passed via the command line.
	fmt.Print(os.Args, "\n")

	// Check the number of command line arguments.
	if len(os.Args) == 1 {
		// Only the path of the executable is provided.
		// Print the path of the executable (index 0).
		fmt.Print(os.Args[0], "\n")
	} else if len(os.Args) > 1 {
		// More than one argument is provided.
		// Print all command line arguments.
		fmt.Print(os.Args, "\n")
	} else {
		// This case will not occur as len(os.Args) is always at least 1.
		fmt.Println("No command line arguments provided.")
	}

	// len returns the length of the slice or string.
	// Here, it prints the length of the os.Args slice and the length of the string "hello".
	fmt.Println(len(os.Args), len("hello"))
}
