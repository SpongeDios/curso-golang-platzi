package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	empleado := Employee{
		id:   1,
		name: "Hector Alvarado",
	}

	empleadoVacio := Employee{}

	fmt.Println(empleado)
	fmt.Println(empleadoVacio)

	empleadoVacio.id = 2
	empleadoVacio.name = "Cristiano Ronaldo"

	fmt.Println(empleadoVacio)
}
