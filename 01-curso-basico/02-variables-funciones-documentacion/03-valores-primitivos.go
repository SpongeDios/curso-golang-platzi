package main

import "fmt"

func main() {
	//Numeros enteros
	var a int
	var b int8
	var c int16
	var d int32
	var e int64

	//Enteros positivos
	var f uint
	var g uint8
	var h uint16
	var i uint32
	var j uint64

	//Numeros decimales
	var k float32
	var l float64

	//TExtos y booleanos
	var n string
	var m bool

	//Numero complejo (numero real e imaginario)
	var o complex64
	var p complex128

	//ejemplo
	complejo := 10 + 8i

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, complejo)
}
