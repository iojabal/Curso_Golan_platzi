package main

import (
	"fmt"
	"prueba/prueba"
)

func main() {

	var Mycar prueba.CarPublic
	Mycar.Brand = "ferrar"
	Mycar.Year = 2000
	fmt.Println(Mycar)

	prueba.PrintMessage("hola platzi")
}
