package main

import "fmt"

func main() {
	printMessage("Hola mundo")
	printMessage("Hola mundo 2")
	printMessage("Hola mundo 3")

	tripleArgument(4, 5, "hola")
	value := functWithReturnValue(10)
	fmt.Println(value)
	value1, value2 := functWithDoubleReturn(10)
	fmt.Println(value1, value2)

	value3, _ := functWithDoubleReturn(50)
	fmt.Println("Ignorando el segundo valor que retorna la funcion", value3)

}

func printMessage(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func functWithReturnValue(a int) int {
	return a * 2
}

func functWithDoubleReturn(a int) (c, d int) {
	return a, a * 2
}
