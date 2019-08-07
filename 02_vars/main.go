package main

import "fmt"

func main() {
	name, email := "Chase", "cmitchellangelo@gmail.com"
	var age = 37
	const isCool = true

	fmt.Println(name, age, email)
	fmt.Printf("%T\n", name)
}
