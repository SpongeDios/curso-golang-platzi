package main

import "fmt"

func main() {
	// DECLARACION DE CONSTASTES
	const pi float32 = 3.14
	const pi2 = 3.1415
	const hola string = "Hola mundo"

	base := 12
	var altura int = 14
	var area int

	println(hola, pi)
	println("pi1: ", pi)
	println("pi2: ", pi2)

	fmt.Println(base, altura, area, pi)

	/*
		OJO: la diferencia en println nativo y el package fmt, es que fmt te entrega valores mejor formateados
	*/

	//Zero values

	var a int     //0
	var b float64 //0.0
	var c string  // ""
	var d bool    // false

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado int = 10
	const areaCuadrado int = baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)

}
