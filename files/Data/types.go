package data

import "fmt"

//creating alias
type integer = int
type sliceOfStrings = []string
type myJson = map[string]string

var x integer = 12

//if we omit the "=" we create a new type. Why ? For instace: semantics
type distance float32
type kilometers float32

// func ToKm(miles distance) kilometers {
// 	return kilometers(miles * 1.60934)
// }

//method
func (miles distance) ToKm() kilometers {
	return kilometers(miles * 1.60934)
}

func (km kilometers) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(4.5)
	km := d.ToKm()
	km.ToMiles()
	fmt.Println(d)
}
