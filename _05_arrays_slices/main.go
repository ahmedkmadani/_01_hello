package main

import "fmt"

func main() {
	// 	//Arrays

	// 	var fruitArray [2]string{"Apple", "Orange"}
	// 	fruitArray := [2]string{"Apple", "Orange"}

	fruitArray := []string{"Apple", "Orange", "Grape"}

	// Assign values
	// fruitArray[0] = "Apples"
	// fruitArray[1] = "Orange"

	fmt.Println(fruitArray)
	fmt.Println(len(fruitArray))
	fmt.Println(fruitArray[1:2])
	fmt.Println(fruitArray[1])

}
