package main

import "fmt"

func main() {

	ids := []int{33, 45, 24, 45, 0}

	//Loop through ids using range

	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with MAP

	emails := map[string]string{"Boob": "boob@mail.com", "Mike": "mike@mail.com"}

	for k, v := range emails {
		fmt.Printf("%s is name - %s is email \n", k, v)
	}

	for k := range emails {
		fmt.Printf("%s name", k)
	}

}
