package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Department string
}

func (m Manager) Work() string {
	return "Name: " + m.Name + "\n" +
		"ID: " + m.ID
}
func main() {
	manager1 := Manager{
		Employee: Employee{
			Name: "Yaroslav Vasilyev",
			ID:   "87665756",
		},
		Department: "Projects",
	}

	fmt.Println(manager1.Work())
}
