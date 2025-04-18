package main

import (
	"fmt"
)

func main() {
	// allNumbers := numbers[:] //slice copy of our array
	// firstThree := numbers[0:3] 
	numbers := []int {10,20,30,40,50}

	fruits := []string{"apple", "banana","strawberry"}
	fmt.Printf("This is my fruits: %v\n", fruits)

	fruits = append(fruits, "kiwi")
	fmt.Printf("This is my new fruits 1 array: %v\n", fruits)

	fruits = append(fruits, "kiwi", "orange", "mango")
	fmt.Printf("This is my new fruits 2 array: %v\n", fruits)

	moreFruits := []string{"grape", "tomato"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("This is my new fruits 3 array: %v\n", fruits)


	for index,value := range numbers{
		fmt.Printf("Index: %d and value %d\n", index,value)
	}
}