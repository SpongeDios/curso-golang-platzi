package main

import "fmt"

func main() {
	//Defer (Antes de morir, esta funcion, se ejecutara la instruccion anotada con defer)
	defer fmt.Println("Hola")
	defer fmt.Println("Aparece arriba de Hola")

	//Continue y break
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

		//contine
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 3 {
			fmt.Println("Break")
			break
		}
	}

	fmt.Println("Mundo")

}
