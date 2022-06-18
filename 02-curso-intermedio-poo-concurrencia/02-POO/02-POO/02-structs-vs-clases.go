package main

import "fmt"

type employee struct {
	id   int
	name string
}

func main() {
	empleado := employee{
		id:   1,
		name: "Hector Alvarado",
	}

	empleadoVacio := employee{}

	fmt.Println(empleado)
	fmt.Println(empleadoVacio)

	empleadoVacio.id = 2
	empleadoVacio.name = "Cristiano Ronaldo"

	fmt.Println(empleadoVacio)
}
