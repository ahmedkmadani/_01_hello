package main

import "fmt"

//Define person struct

type Person struct {
	firstName string
	lastname  string
	city      string
	gender    string
	age       int
}

//Greeting Method (values receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastname
}

func main() {
	person1 := Person{firstName: "Ahmed", lastname: "Kamal", city: "LA", gender: "M", age: 25}
	person2 := Person{"Ahmed", "Kamal", "LA", "M", 25}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.greet())
}
