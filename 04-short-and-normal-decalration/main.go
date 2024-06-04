package main

import "fmt"

// package scoped variable use normal decarlation not use short
// version := "0"
var version string

func main() {
	// if you don't know the initial value
	// don't use short declartion
	// var score = 0
	// this already take default value as 0
	var score int

	// if we know initial value then use short decarlation
	name := "Zehan"

	height, length := 10, 20

	// don't use this
	// height = 20
	// color := "pink"

	// use this
	height, color := 20, "pink"

	// group variable for readablity
	var (
		video     string
		timeStamp string
		duration  int
	)

	fmt.Println(score, name, height, color, length, video, timeStamp, duration)
}
