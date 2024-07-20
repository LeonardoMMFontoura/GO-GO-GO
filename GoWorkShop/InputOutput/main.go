package main

import "leonardomastra.dev/data"
import "fmt"

func main() {
	message := "Hello World from the main file!\n"
	print(message)
	fmt.Println("Hello World from the main file2!")
	fmt.Println(data.Price)
	printToScreen()
	
	result := add(1,2)
	fmt.Println("result add" , result)
	resultadd , _ := addAndSubtract(1,2)
	fmt.Println("resultado ", resultadd)

	age := 29
	addAge(&age)
	fmt.Println("Should return: 30, returned",age)
}


func add(x int, y int) int {
	return x + y
}

func addAndSubtract(x int, y int) (int, int) {
	return x + y, x - y
}

func addAge(age *int) {
	*age++
}