package main

import "fmt"

func main(){
	fmt.Println(greet("Samuelson", "Okoi"))
	fmt.Println("Hello,", greet("Samuelson", "Okoi"))
	fmt.Println(hello("Samuelson", "Okoi"))
	fmt.Println(rmultiple("Samuelson", "Okoi"))
	fmt.Println(average(43, 56, 87, 12, 45, 57))

	data := []float64 {43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)

	// func expression
	sayHi := func(){
		fmt.Println("Hello there.")
	}

	sayHi()

	fmt.Printf("%T\n", sayHi)
}

// returning a string
func greet(fname string, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}

// returning a named string
func hello(fname, lname string) (s string) {
	s = fmt.Sprint("Hi! ", fname, " ", lname)
	return
}

// returning multiple values
func rmultiple(fname, lname string) (string, string){
	return fmt.Sprint(fname, " ", lname), fmt.Sprint(lname, " ", fname)
}

// Variadic funtions parameter
func average(sf ...float64) float64 {
	fmt.Printf("%T \n", sf)
	fmt.Println(sf)
	var total float64
	for _, s := range sf {
		total += s
	}
	return total / float64(len(sf))
}