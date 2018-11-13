package main

import "fmt"

func main(){
	var name = ""

	fmt.Print("What is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello, ", name)
}