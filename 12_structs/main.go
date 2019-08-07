package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type person struct {
	/*	firstName string
		lastName  string
		city      string
		gender    string
		age       int
	*/
	firstName, lastName, city, gender string
	age                               int
}

// Gretting method (value reciever)
func (p person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " from " + p.city
}

// hasBirthday method (pointer reciever)

func (p *person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)

func (p *person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// Init person using struct
	person1 := person{firstName: "Chase", lastName: "Mitchell", city: "Austin", gender: "f", age: 25}

	// Alternative
	// person2 := Person{"Chase", "Mitchell", "Austin", "m", 25}

	fmt.Println(person1.firstName)
	// person1.age++
	fmt.Println(person1)

	//fmt.Println(person2)
	//person2.age++
	//fmt.Println(person2.firstName)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

}
