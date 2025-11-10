package main

import "fmt"

func main() {

	var x [5]int
	fmt.Println(x)
	x[0] = 42
	x[0] = 34
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	var ni [7]int
	ni[0] = 33
	fmt.Printf("%v \t\t\t  %T\n ", ni, ni)

	n := [4]int{55, 54, 57, 58}
	fmt.Printf("%v \t\t\t  %T\n ", n, n)

	ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Printf("%v \t\t\t  %T\n ", ns, ns)

	fmt.Println(cap(x))
	fmt.Println(len(ni))
	fmt.Println(cap(ni))
	fmt.Println(len(n))
	fmt.Println(len(x))
	fmt.Println(cap(n))
	fmt.Println(len(ns))
	fmt.Println(cap(ns))

	xs := []string{"hello", "world", "Gophers"}
	fmt.Println(xs)

	xm := []string{"Mumbai", "Delhi", "Bangalore", "Hyderabad", "Chennai", "Pune", "Kolkata", "Ahmedabad", "Jaipur", "Surat", "Lucknow", "Kanpur", "Nagpur", "Indore", "Thane", "Bhopal", "Visakhapatnam", "Vadodara", "Patna", "Ludhiana"}
	fmt.Printf("xm %v\n", xm)

	op := [5]int{}
	for i := 0; i < 5; i++ {
		op[i] = i
	}

	for i, v := range op {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
}
