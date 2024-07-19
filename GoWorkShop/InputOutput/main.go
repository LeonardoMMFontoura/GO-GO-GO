package main

import "leonardomastra.dev/data"
import "fmt"

func main() {
	message := "Hello World from the main file!\n"
	print(message)
	fmt.Println("Hello World from the main file!")
	fmt.Println(data.Price)
	printToScreen()
}
