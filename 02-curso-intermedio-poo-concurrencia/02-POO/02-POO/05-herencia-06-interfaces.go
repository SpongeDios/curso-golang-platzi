package main

import "fmt"

type Empleadoo struct {
	name     string
	vacation bool
}

type Persona struct {
	id  int
	age int
}

type FullTimeEmpleadoo struct {
	Persona
	Empleadoo
	endDate string
}

func (fullTimeEmpleadoo FullTimeEmpleadoo) getMessage() string {
	return "Este es un FullTimeEmpleadoo"
}

type PrintInfo interface {
	getMessage() string
}

type EmpleadooTemporal struct {
	Persona
	Empleadoo
	taxRate int
}

func (empleadoTemporal EmpleadooTemporal) getMessage() string {
	return "Este es un empleado temporal"
}

func getMessage(info PrintInfo) {
	fmt.Println(info.getMessage())
}

func main() {
	fullTimeEmpleado := FullTimeEmpleadoo{}
	fullTimeEmpleado.name = "Nombre"
	fullTimeEmpleado.age = 26
	fullTimeEmpleado.id = 1
	fullTimeEmpleado.vacation = true

	fmt.Println(fullTimeEmpleado)
	//Go no tiene una herencia como tal, pero esto se soluciona rapidamente con interfaces

	empleadooTemporal := EmpleadooTemporal{}
	getMessage(fullTimeEmpleado)
	getMessage(empleadooTemporal)
}
