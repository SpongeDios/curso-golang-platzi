package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) Id() int {
	return e.id
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func main() {
	nuevoEmpleado := Employee{}
	fmt.Println(nuevoEmpleado)
	nuevoEmpleado.SetId(3)
	nuevoEmpleado.SetName("Hector Alvarado")
	fmt.Println(nuevoEmpleado)
	fmt.Println(nuevoEmpleado.id, nuevoEmpleado.name, nuevoEmpleado.Id(), nuevoEmpleado.Name())
}
