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

	Tesla := struct {
		CEO                  string
		Revenue              float32
		manufacturingProcess string
	}{
		CEO:                  "Elon Musk",
		Revenue:              36,
		manufacturingProcess: "who knows what happens inside tesla gigafactories!",
	}
	fmt.Println(Tesla.CEO)
	fmt.Println(Tesla.Revenue, "billion dollars")
	fmt.Println(Tesla.manufacturingProcess)
}
