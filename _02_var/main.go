package main

import "fmt"

func main() {

	//Short hand var declartion
	// name := "Ahmed"
	// email := "brand@mail.com"

	name, email := "Ahmed", "brand@mail.com"
	size := 9.77
	var age = 90

	const isCool = true

	fmt.Println(name, age, size, email)

	//get var type using printf and %T
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", isCool)

}
