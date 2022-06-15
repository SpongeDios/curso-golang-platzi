package main

import "fmt"

func main() {
	//PRIMERA FORMA
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//SEGUNDA FORMA

	otroModulo := 5 % 2

	switch otroModulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//SIN CONDICION
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No cumple ninguna condicion")
	}
}
