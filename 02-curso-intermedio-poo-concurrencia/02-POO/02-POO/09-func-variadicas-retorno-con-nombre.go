package main

import "fmt"

func sum2(a, b int) int {
	return a + b
}

func sum(a ...int) int {
	var suma int
	for _, value := range a {
		suma += value
	}
	return suma
}

func printNames(args ...string) {
	for _, values := range args {
		fmt.Println(values)
	}
}

func getValuesFeo(x int) (int, int, int) {
	return x * 2, x * 3, x * 4
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum2(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))
	printNames("Hector", "German", "ALvarado", "Campos")
	fmt.Println(getValuesFeo(2))
	fmt.Println(getValues(2))
}
