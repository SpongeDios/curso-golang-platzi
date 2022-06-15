package main

import (
	"04-estructuras-datos-basicas/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.SetBrand("Ferrari")
	myCar.SetYear(2022)
	fmt.Println(myCar)

	mypackage.PrintMessage("Hola mundo")
}
