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
	x := 42
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
	/* switch = scd (switch , case , default)
	 */
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
	switch x {
	case 40:
		fmt.Println("x is 40", x)
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrought: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthrought x is 42")
		fallthrough

	case 43:
		fmt.Println("printed because of fallthrough x is 43")
		fallthrough
	default:
		fmt.Println("printed because of ALL The fallthrough statements: this is the default case for x")
	}

	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	/*m := map[string]int{
		"James": 99,
		"Pooja": 98,
	}*/
	/*for k, v := range m {
		fmt.Println("Ranging over a map", k, v)
	}
	*/
	fmt.Println("-------------------------------------------")
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v \n and the total count is %v and x is %v ", i, c, x)
		}
	}
	fmt.Println("-------------------------------------------")
	m := map[string]int{
		"James":       42,
		"Money penny": 32,
	}
	for k, v := range m {
		fmt.Printf("Key %v \t value is%v\n", k, v)
	}
	age1 := m["James"]
	fmt.Printf("The age is %v\n", age1)

	age2 := m["Q"]
	fmt.Printf("The age is %v\n", age2)

	if v, ok := m["James"]; !ok {
		fmt.Printf("The age is %v\n", v)
	}
}
