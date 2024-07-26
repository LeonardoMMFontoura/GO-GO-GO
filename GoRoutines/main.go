package main

import "fmt"

func PrintMessage(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
	}
}
func main() {
	PrintMessage("Go is fucking amazing!")
	PrintMessage("And I love it!")
}
