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

}
