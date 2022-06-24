package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}

func isNil(err error, operador int) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(operador)
	}
}

type calc struct {
}

func (c calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parseador(entradaLimpia[0])
	operador2 := parseador(entradaLimpia[1])

	outputNumber := 0
	switch operador {
	case "+":
		outputNumber = operador1 + operador2
		break
	case "-":
		outputNumber = operador1 - operador2
		break
	case "*":
		outputNumber = operador1 * operador2
		break
	case "/":
		outputNumber = operador1 / operador2
		break
	default:
		fmt.Printf("El operador %s no esta soportado en esta calculadora.", operador)
	}
	return outputNumber
}

func parseador(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
