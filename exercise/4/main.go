package main

import "fmt"

func main(){
	var smallnum, bignum int

	fmt.Print("Enter a small number: ")
	fmt.Scanf("%d", &smallnum)

	fmt.Println("Enter a bigger number: ")
	fmt.Scanf("%d", &bignum)

	fmt.Println(bignum, "%", smallnum, "=", bignum % smallnum)
}