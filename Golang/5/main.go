package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  string
}

func (p Person) Greet() string {
	return "Hi! My name is " + p.Name + " and I'm " + p.Age + " years old"
}
func main() {
	person1 := Person{
		Name: "Yaroslav Vasilyev",
		Age:  "20",
	}

	fmt.Println(person1.Greet())
}
