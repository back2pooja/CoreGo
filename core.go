package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("go core started....")
		fmt.Println("go version: " + runtime.Version())

	}
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 50
	y := 5

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
		fmt.Println("x is getting close to the meaning of life")
	}
	if x != 42 && y != 10 {
		fmt.Println("x is not 42 and y is not 10")
	}

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and GREATER THAN OR EQUALS x which is %v\n\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n\n", z, x)
	}
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN ", x)
	case x == 42:
		fmt.Println("x is EQUAL to 42 ", x)
	case x > 42:
		fmt.Println("x is GREATER THAN ", x)
	default:
		fmt.Println("x is default case for x ", x)
	}

	switch x {
	case 40:
		fmt.Println("x is LESS THAN 40 ", x)
	case 41:
		fmt.Println("x is GREATER THAN 41", x)
	case 42:
		fmt.Println("x is EQUAL to 42 ", x)
	case 43:
		fmt.Println("x is EQUAL to 43  ", x)
	default:
		fmt.Println("x is default case for x ", x)

	}
}
