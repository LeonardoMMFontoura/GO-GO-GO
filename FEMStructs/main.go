package main

import (
	"fem/structs/model"
	"fmt"
)

func main() {
	// leo := model.Instructor{Id: 21, Firstname: "leonardo", Lastname: "mastra", Score: 32}
	kyle := model.NewInstructor("Leonardo", "Mastra")
	print(kyle.Print())
	goCourse := model.Course{Id: 2, Name: "Leo Learning", Instructor: kyle}
	fmt.Printf("%v", goCourse)
	// print(leo.Print())
	// fmt.Println(leo)
	// leo.Firstname = "Leonardo"
}
