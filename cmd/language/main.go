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

type Company struct {
	CEO               string
	Revenue           float32
	NumberOfEmployees int
	AveragePay        float32
	secretSauce       string
}

func (w *Water) printWater() {
	fmt.Println("\tBrand:", w.brand)
	fmt.Println("\tPrice (MMK):", w.price)
	fmt.Println("\tCapacity (in L):", w.capacity_in_litres)
}

func (c *Company) changeCEO(name string) {
	c.CEO = name
}

func main() {
	Alpine := Water{
		brand:              "Alpine",
		price:              300,
		capacity_in_litres: 1,
	}
	Alpine.printWater()
	fmt.Println()

	twitter := Company{
		CEO:               "Jack Dorsey",
		Revenue:           4,
		secretSauce:       "None",
		NumberOfEmployees: 2000,
		AveragePay:        100000,
	}
	fmt.Println(twitter)
	twitter.changeCEO("Elon Musk")
	fmt.Println(twitter.CEO)
	fmt.Println(twitter)
	renumerationInMillions := twitter.calculateTotalAnnualEmployeeCompensation() / 1000000
	fmt.Println(renumerationInMillions, "million dollars")
}

func (c Company) calculateTotalAnnualEmployeeCompensation() float32 {
	return float32(c.NumberOfEmployees) * c.AveragePay
}
