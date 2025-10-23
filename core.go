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
	x := 42
	y := 42
	fmt.Printf("x=%v, y=%v\n", x, y)
	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}
	/*
	   "if " statement specify the conditional execution of two branches according tp the value of a boolean expression. if expression evaluates to true
	   , the "if" branch is executed , otherwise if present , the "else" branch is executed/

	*/
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
	/*
		    Logical operators apply to boolean values
		     and yield a result of thr same type as the operands.
		the right oprand is evaluated conditionally
	*/

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is greater than the meaning of life")
	}
	if x != 42 && y != 10 {
		fmt.Println("x is not 42 and y is not 10")
	}

}
