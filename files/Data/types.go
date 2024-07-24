package data

import "fmt"

//creating alias
type integer = int
type sliceOfStrings = []string
type myJson = map[string]string

var x integer = 12

//if we omit the "=" we create a new type. Why ? For instace: semantics
type distance float32

func test() {
	d := distance(4.5)
	fmt.Println(d)
}
