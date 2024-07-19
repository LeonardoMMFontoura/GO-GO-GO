package data

import "fmt"

var Countries [10]string = [10]string{"Brazil", "USA", "Canada", "Mexico", "Argentina", "Chile", "Peru", "Colombia", "Uruguay", "Paraguay"}
var Slice []int
var Codes map[string]int


func init() {
	WellKnowPorts := map[string]int {"http": 80, "https": 443, "ftp": 21, "ssh": 22}
	Countries[0] = "Brazill"
	Prices := [5]int{1, 2, 3, 4, 5}
	names := []string {"Leonardo", "Mastra", "Dev", "Duda"}
	names = append(names, "Jorge")
	names = append(names, "Mariza")
	quantityOfCountries := len(Countries)
	fmt.Println("Well Know Ports:", WellKnowPorts)
	fmt.Println("Quantity of countries:", quantityOfCountries)
	fmt.Println("Prices:", Prices)
	fmt.Println("Names:", names)
}