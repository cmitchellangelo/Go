package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { // I can just have int set once, do not need both
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Chase"))
	fmt.Println(getSum(3, 7))
}
