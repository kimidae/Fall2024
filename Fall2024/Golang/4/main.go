package main

import "fmt"

func add(a int, b int) int {
	var c int = a + b
	return c
}
func swap(a int, b int) (int, int) {
	return b, a
}
func qwre(a int, b int) (int, int) {
	var c int = a / b
	var d int = a % b
	return c, d
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(add(a, b))
	fmt.Println(swap(a, b))
	fmt.Println(qwre(a, b))
}
