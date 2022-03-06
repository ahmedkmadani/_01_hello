package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short Method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//FIZZBUZZ

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FIZZBUZZ")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}
