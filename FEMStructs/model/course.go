package model

import "fmt"

type Duration float32

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) string() string {
	return fmt.Sprintf("--- %v --- (%v)\n", c.Name, c.Instructor.Firstname)
}
