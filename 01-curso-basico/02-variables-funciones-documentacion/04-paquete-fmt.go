package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	//v se utiliza para cualquier valor
	//la buena practica es utilizar el tipo de dato
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//Sprintf (para guardar la print en una variable :O)
	message := fmt.Sprintf("%s tiene mas de %d cursos ", nombre, cursos)
	fmt.Println(message)

	//Conocer el tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)

}
