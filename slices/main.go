package main

import "fmt"

func main() {
	// Slice declaration
	days := []string { "Monday", "Tuesday", "Wednessday" }

	fmt.Println("Content of the list: ", days)

	for index, day := range days {
		fmt.Println(index, day)
	}

	days = append(days, "Thursday", "Friday")

	fmt.Println("Appended some new values to the slice")

	for index, day := range days {
		fmt.Println(index, day)
	}

}