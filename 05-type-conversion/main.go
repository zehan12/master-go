package main

import "fmt"

func main() {
	force := 100
	acceleration := 2.7
	mass := 200
	force = mass * int(acceleration) // 200 * 2 = 400
	fmt.Println(force)

	force = int(float64(mass) * acceleration) // 200.0 * 2.7 = 540
	fmt.Println(force)

}
