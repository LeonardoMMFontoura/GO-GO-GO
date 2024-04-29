package main

import "fmt"

func define() string {
	const s string = "leonardo"
	return s
}

func main() {
	var pega = define()
	fmt.Println(pega)
}
