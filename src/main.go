package main

import "fmt"

type fig2d interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type Rectangle struct {
	base   float64
	altura float64
}

func (r Rectangle) area() float64 {
	return r.altura * r.base
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func Calculate(f fig2d) {
	fmt.Println("el area es: ", f.area())
}

func main() {
	MySquare := cuadrado{base: 10}
	MyRectangle := Rectangle{base: 2, altura: 4}

	Calculate(MySquare)
	Calculate(MyRectangle)

	//lista de interfaces
	myInterface := []interface{}{"hola", 12, 0.123}
	fmt.Println(myInterface...)

}
