package main

import "fmt"

func main() {
	myCar := car{
		brand: "Chevrolet",
		year:  2022,
		color: "white",
		model: "S",
		full:  true,
	}

	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2022
	otherCar.color = "Red"
	otherCar.model = "S"
	otherCar.full = true

	fmt.Println(myCar)
	fmt.Println(otherCar)

}

type car struct {
	brand string
	year  int
	color string
	model string
	full  bool
}
