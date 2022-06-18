package main

import "fmt"

type Empleado struct {
	id       int
	name     string
	vacation bool
}

func NewEmpleado(id int, name string, vacation bool) *Empleado {
	return &Empleado{id: id, name: name, vacation: vacation}
}

func main() {
	//1
	empleado := Empleado{}
	fmt.Println(empleado)
	//2
	empleado2 := Empleado{
		id:       2,
		name:     "Hector Alvarado",
		vacation: true,
	}
	fmt.Println(empleado2)
	//3
	empleado3 := new(Empleado) //devuelve una referencia
	fmt.Println(*empleado3)

	empleado4 := NewEmpleado(4, "Hectorrrrr", true)
	fmt.Println(*empleado4)
}
