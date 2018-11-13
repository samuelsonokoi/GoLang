package main

import "fmt"

func main(){
	a := 50

	fmt.Println("variable a holds ", a)
	fmt.Println("memory address of variable a is ", &a)

	var b *int = &a

	fmt.Println("referencing a pointer ", b)
	fmt.Println("dereferencing a pointer ", *b)

	*b = 100

	fmt.Println("new value of a is ", a)


}