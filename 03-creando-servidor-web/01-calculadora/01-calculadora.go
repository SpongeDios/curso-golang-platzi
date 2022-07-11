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

func (c calc) operate(entrada string, operador string) (int, error) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1, err := parseador(entradaLimpia[0])

	if err != nil {
		return -1, err
	}
	operador2, err2 := parseador(entradaLimpia[1])

	if err2 != nil {
		return -1, err2
	}

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
	return outputNumber, nil
}

func parseador(entrada string) (int, error) {
	operador, err := strconv.Atoi(entrada)
	if err != nil {
		return -1, fmt.Errorf("error formating string to int %v", err)
	}
	return operador, nil
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
