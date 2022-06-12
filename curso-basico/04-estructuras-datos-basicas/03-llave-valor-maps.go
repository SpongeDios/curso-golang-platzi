package main

import "fmt"

func main() {
	/*
	Los mapas son mas eficientes que manejar arrays o slices puesto que
	de manera nativa implementan concurrencia
	*/
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	//recorrer mapa
	//TODO: LOS MAPAS SON DESORDENADOS OJO, si queremos orden tenemos los slices
	for clave, valor := range m {
		fmt.Println(clave, valor)
	}

	//encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	//si no existe la llave, devuelve un zero value
	value = m["no_existe"]
	fmt.Println(value)

	//Verificar si realmente existe la llave con un valor
	valor, ok := m["Jose"]
	fmt.Println(valor, ok)
}

