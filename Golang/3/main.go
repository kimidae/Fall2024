package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a == 0 {
		fmt.Println("Zero")
	} else if a > 0 {
		fmt.Println("Positive")
	} else if a < 0 {
		fmt.Println("Negative")
	}
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 10 first natural number is: ", sum)
	fmt.Print("Enter number from 1 to 7:\n")
	fmt.Scan(&a)
	switch {
	case a == 1:
		fmt.Println("Monday")
	case a == 2:
		fmt.Println("Tuesday")
	case a == 3:
		fmt.Println("Wednesday")
	case a == 4:
		fmt.Println("Thursday")
	case a == 5:
		fmt.Println("Friday")
	case a == 6:
		fmt.Println("Saturday")
	case a == 7:
		fmt.Println("Sunday")
	}

}
