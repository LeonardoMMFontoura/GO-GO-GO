package main

import (
	"fmt"
)

func main() {

	// var name string = "leonardo"
	// fmt.Printf("My name is %s\n", name)
	// age:= 27
	// fmt.Printf("I'm %d\n", age)
	// var city string
	// city = "Rio de Janeiro"
	// fmt.Printf("%s", city)

	var city, continent, name string
	city = "Rio de janeiro"
	continent = "America"
	name = "teste"
	fmt.Printf("%s %s %s\n", city, continent, name)

	var (
		isEmployed bool = true
		salary int = 15.000
		positon string = "developer"
	)
	fmt.Printf("IsEmployed: %t , This is my salary: %d, and my position: %s\n", isEmployed, salary, positon)

	//zero values
	var defaultInt int 
	var defaultBool bool
	var defaultFloat float64
	var defaultString string
	fmt.Printf("valor int %d, valor bool %t, valor float %f, valor string %s", defaultInt, defaultBool, defaultFloat, defaultString)

	const pi = 3.14
	const (
		Monday = 1
		Tuesday = 2
		Wednesday = 3
	)
	fmt.Printf("Monday - %d, Tuesday %d - Wednesday - %d\n", Monday, Tuesday, Wednesday)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
	)

	fmt.Printf("Jan - %d Feb - %d Mar - %d Apr - %d", Jan,Feb, Mar, Apr)


	result := add(3,4)
	fmt.Printf("Result = %d", result)
}

func add (a int, b int) int {
	return a + b
}