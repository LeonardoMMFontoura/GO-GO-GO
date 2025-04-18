package main

import (
	"fmt"
)

func main() {
	numbers := [5]int {10,20,30,40,50} // arrays are imutable once declared, but we can subscribe the elements at index like above
	numbers[0] = 100 // but we cant grow or shrink the array
	fmt.Println("This is my array %v", numbers)

	matrix := [2][3]int{
		{1,2,3},
		{4,5,6},
	}
}