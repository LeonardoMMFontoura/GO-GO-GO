package main

import (
	"fem/structs/model"
)

func main() {
	leo := model.Instructor{Id: 21, Firstname: "leonardo", Lastname: "mastra", Score: 32}
	print(leo.Print())
	// fmt.Println(leo)
	// leo.Firstname = "Leonardo"
}
