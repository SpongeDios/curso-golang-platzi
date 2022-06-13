package main

import "fmt"

func main() {
	myCuadrado := cuadrado{base: 2}
	area := myCuadrado.area()
	fmt.Println(area)

	myRectangulo := rectangulo{base: 2, altura: 4}
	areaRectangulo := myRectangulo.area()
	fmt.Println(areaRectangulo)

	myTriangulo := triangulo{base: 4, altura: 8}
	//ejecutando la interface
	calcular(myCuadrado)
	calcular(myRectangulo)
	calcular(myTriangulo)

	//LISTA INTERFACES
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface)
	//Los tres puntos imprime cada uno por separado
	fmt.Println(myInterface...)
}

/////////CUADRADOOOOO///////
type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

//////RECTANGULO///////

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

//////TRIANGULO//////

type triangulo struct {
	base   float64
	altura float64
}

func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

//////INTERFACEEEEEEE///////

type figuras2D interface {
	area() float64
}

func calcular(f figuras2D) {
	fmt.Println("Area", f.area())
}
