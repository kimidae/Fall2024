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

func ToJSON(p Product) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FromJSON(jsonStr string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func main() {
	product := Product{
		Name:     "Laptop",
		Price:    999.99,
		Quantity: 10,
	}

	jsonStr, err := ToJSON(product)
	if err != nil {
		fmt.Println("Error encoding product:", err)
		return
	}
	fmt.Println("JSON representation:", jsonStr)

	decodedProduct, err := FromJSON(jsonStr)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}
