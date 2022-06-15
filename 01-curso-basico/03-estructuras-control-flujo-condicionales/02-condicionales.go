package main

import "fmt"

const usuario string = "hector"
const contrasena string = "12345678"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad (AND)")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("La segunda condicion es verdadera (OR)")
	}

	parImpar(10)
	isLogged := login("hector", "12345678")
	fmt.Println(isLogged)

}

//func es par o impar
func parImpar(number int) {
	if number%2 == 0 {
		fmt.Println("Es par")
		return
	}

	fmt.Println("Es impar")
}

func login(username, password string) bool {
	if username == usuario && password == contrasena {
		return true
	}
	return false
}

//func usuario y password es correcto return true (LOGIN)
