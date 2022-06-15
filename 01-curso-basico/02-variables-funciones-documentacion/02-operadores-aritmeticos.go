package main

import "fmt"

func main() {

	//Area cuadrado
	const baseCuadrado int = 10
	const areaCuadrado int = baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado:", areaCuadrado)

	var x int = 10
	var y int = 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Calcula el area de un trapecio rectangulo y circulo
	var ladoLargo int = 10
	var ladoCorto int = 5

	result = ladoCorto * ladoLargo
	fmt.Println("Area de un rectangulo:", result)

	var pi float32 = 3.14
	var radio float32 = 15

	floatResult := (radio * radio) * pi
	fmt.Println("Area de un cuadrado:", floatResult)

	//Area de un trapecio
	ladoArriba := 3
	ladoAbajo := 7
	altura := 4

	floatResult = ((float32(ladoArriba) + float32(ladoAbajo)) / 2) * float32(altura)
	fmt.Println("Area de un trapecio:", floatResult)

}
