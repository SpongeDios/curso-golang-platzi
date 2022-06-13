package main

import "fmt"

func main() {
	//El simbolo & hace referencia al espacio en memoria, ejemplo: #123AXBJ1
	//El * es el simbolo contrario de &: El * trae el valor de un espacio en memoria
	//Ejemplo de puntero: #123AXBJ1: "Hola mundo"
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()
	//myPc.ram = 20
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()

	fmt.Println(myPc)

	p := newPc(12, 1024, "ojo")
	fmt.Println(p)
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func newPc(ram int, disk int, brand string) *pc {
	return &pc{ram: ram, disk: disk, brand: brand}
}

//EL ASTERISCO QUE SE ANTEPONE EN PC ES PARA DECIRLE QUE ACCEDEREMOS A LOS VALORES
//DE myPc MEDIANTE EL PUNTERO

func (myPc *pc) Ram() int {
	return myPc.ram
}

func (myPc *pc) SetRam(ram int) {
	myPc.ram = ram
}

func (myPc *pc) Disk() int {
	return myPc.disk
}

func (myPc *pc) SetDisk(disk int) {
	myPc.disk = disk
}

func (myPc *pc) Brand() string {
	return myPc.brand
}

func (myPc *pc) SetBrand(brand string) {
	myPc.brand = brand
}

func (myPc *pc) ping() {
	fmt.Println(myPc.disk)
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}
