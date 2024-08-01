package main

import "fmt"

func PrintMessage(message string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
	}
	channel <- "Done!"
}
func main() {
	var channel chan string
	go PrintMessage("Go is fucking amazing!", channel)
	// PrintMessage("And I love it!")
	println(<-channel)
}
