package main

import (
	"fmt"
	"reflect"
)

func main() {

	tst := "string"
	tst2 := 10
	tst3 := 1.2
	tst4 := true
	tst5 := []int{1, 2, 3}
	tst6 := [3]int{1, 2, 3}

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))
	fmt.Println(reflect.TypeOf(tst4))
	fmt.Println(reflect.TypeOf(tst5), "slice")
	fmt.Println(reflect.TypeOf(tst6), "array")

}
