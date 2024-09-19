package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func Write(p Product) string {
	data, _ := json.Marshal(p)
	return string(data)
}

func Read(res_str string) Product {
	var p Product
	json.Unmarshal([]byte(res_str), &p)
	return p
}

func main() {
	product := Product{
		Name:     "Morozhenoe Raduga",
		Price:    230,
		Quantity: 15.00,
	}

	res_str := Write(product)
	fmt.Println("As JSON:", res_str)

	res_decoded := Read(res_str)
	fmt.Printf("Decoded: %+v\n", res_decoded)
}
