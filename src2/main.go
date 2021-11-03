package main

import "fmt"

func main() {
	//declaracion de variables
	helloMessage := "hello"
	worldMessage := "world"

	//println
	fmt.Println(helloMessage, worldMessage)

	//printf
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tine mas de %d cursos\n", nombre, cursos)
	//si no se sabe lo que ira en el % se le agrega una v

	//Sprintf
	message := fmt.Sprintf("%v tine mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	//tipo de datos

	fmt.Printf("helloMessage: %T", helloMessage)

}
