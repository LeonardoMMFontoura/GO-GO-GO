package main

import "leonardomastra.dev/data"
import "fmt"

func main() {
	message := "Hello World from the main file!\n"
	print(message)
	fmt.Println("Hello World from the main file2!")
	fmt.Println(data.Price)
	printToScreen()
}


func add(x int, y int) int {
	return x + y
}

func addAndSubtract(x int, y int) (int, int) {
	return x + y, x - y
}