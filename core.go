package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("go core started....")
		fmt.Println("go version: " + runtime.Version())

	}
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40
	y := 5
	fmt.Printf("x=%v, y=%v\n", x, y)
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equals to, or greater than, the meaning of life")
	}
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equals to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

}
