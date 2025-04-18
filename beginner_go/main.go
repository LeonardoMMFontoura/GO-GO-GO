package main

import (
	"fmt"
)
func main(){
	capitalCities := map[string]string{
		"USA": "Washington D.C.",
		"India": "New Delhi",
		"UK": "London",
	}
	fmt.Printf(capitalCities["USA"])

	capital, exists := capitalCities["Germany"]
	if (exists){
		fmt.Printf("It is here", capital)
	}else{
		fmt.Printf("It is not here")
	}

	delete(capitalCities, "UK")
}