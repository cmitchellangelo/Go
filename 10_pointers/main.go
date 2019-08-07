package main

import "fmt"

func main() {

	// &a points to value of a
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
