package model

import "fmt"

type Instructor struct {
	Id        int
	Firstname string
	Lastname  string
	Score     int
}

func NewInstructor(name string, lastName string) Instructor {
	return Instructor{Firstname: name, Lastname: lastName}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%d %v %v (%d)", i.Id, i.Firstname, i.Lastname, i.Score)
}
