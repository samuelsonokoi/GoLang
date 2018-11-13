package main

import "fmt"

func main(){
	// Print out all the even numbers from 0 to 100
	for i := 0; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "is an even number")
		}
	}
}