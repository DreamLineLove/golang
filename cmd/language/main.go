package main

import "fmt"

type Water struct {
	brand              string
	price              float32
	capacity_in_litres float32
}

type FastFoodChain struct {
	BrandName      string
	Revenue        float32
	NumberOfChains int
	secretFormula  string
}

func main() {
	Alpine := Water{
		brand:              "Alpine",
		price:              300,
		capacity_in_litres: 1,
	}
	fmt.Println(Alpine.brand)
	fmt.Println(Alpine.price)
	fmt.Println(Alpine.capacity_in_litres)

	fmt.Println("Hello World")

}
