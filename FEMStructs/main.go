package main

import (
	"fem/structs/model"
)

func main() {
	// leo := model.Instructor{Id: 21, Firstname: "leonardo", Lastname: "mastra", Score: 32}
	kyle := model.NewInstructor("Leonardo", "Mastra")
	print(kyle.Print())
	// print(leo.Print())
	// fmt.Println(leo)
	// leo.Firstname = "Leonardo"
}
