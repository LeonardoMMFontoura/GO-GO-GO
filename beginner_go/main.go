package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}


func main(){
	person := Person{name: "John", age: 29}
	fmt.Printf("this is my new person: %+v", person)


	employee := struct {
		name string
		age int  
	}{
		"Leo",
		29,
	}
	fmt.Println("This is my brand new employee", employee)
    type Address struct {
		Street string
		City string
	}

	type Contact struct {
		Name string
		Address Address
		Phone string
	}

	contact := Contact{
		Name:  "Mike",
		Address: Address{
			Street: "123 Main Street",
			City: "Anytown",
		},
	}

	fmt.Println("this is my contact: ", contact)
	fmt.Println("this is my employee", employee)

	fmt.Printf("Person name before: %s\n", person.name)
	modifyPerson(&person)
	fmt.Printf("Person name after: %s\n", person.name)


	x := 20
	ptr := &x
	fmt.Printf("value of x: %d and pointer of x: %p\n",x, ptr)
	*ptr = 30
	fmt.Printf("value of x: %d and pointer of x: %p\n",x, ptr)
}

func modifyPerson(person *Person){
	person.name = "Leonardo"
	fmt.Printf("new name: %s\n", person.name)
}

