package main

import "fmt"

func main() {

	//Define a map
	// emails := make(map[string]string)

	//Assign kv
	// emails["Boob"] = "boob@mail.com"
	// emails["Ahmed"] = "ahmed@mail.com"

	//Declar MAP AND KV

	emails := map[string]string{"Boob": "boob@mail.com", "Mike": "mike@mail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Boob"])

	//Delete from MAP

	delete(emails, "Boob")
	fmt.Println(emails)

}
