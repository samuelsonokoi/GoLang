package main

import "fmt"

const first_name = "Samuelson"

const (
	website = "www.samuelsonokoi.com"
	email = "hello@samuelsonokoi.com"
)

func main(){
	const last_name string = "Okoi"

	fmt.Println("Your name is " + first_name + " " + last_name)
}