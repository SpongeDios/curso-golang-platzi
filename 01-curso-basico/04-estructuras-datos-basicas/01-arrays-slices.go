package main

import "fmt"

func main() {
	//TODO: ARRAYS Y SLICES SON ORDENADOS
	/*
	Los mapas son mas eficientes que manejar arrays o slices puesto que
	de manera nativa implementan concurrencia
	*/

	//array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	//CAP = CAPACIDAD MAXIMA DE UN ARRAY
	//LEN = LARGO DEL ARRAY

	//Slice
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//APPEND (SOLO VALIDO PARA slICE, LOS ARRAYS SON INMUTABLES)
	slice = append(slice, 7)
	fmt.Println(slice)

	//APPEND nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
