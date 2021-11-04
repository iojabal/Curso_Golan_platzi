package main

import "fmt"

type Car struct {
	brand string
	year  int
}

func main() {
	myCar := Car{
		brand: "Ford",
		year:  2020,
	}
	fmt.Println(myCar)

	//otra manera
	otherCar := Car{brand: "Toyota", year: 1190}
	otherCar.brand = "Honda"

	fmt.Println(otherCar)
}
