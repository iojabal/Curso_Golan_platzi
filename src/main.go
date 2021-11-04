package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.149516
	const p2 = 3.149616
	fmt.Println("pi", pi)
	fmt.Println("pi2:", p2)
	//declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)
	//zero values
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("el area del cuadrado es igual a: ", areaCuadrado)

}
