package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("go core started....")
	fmt.Printf("go version" + runtime.Version())
}
