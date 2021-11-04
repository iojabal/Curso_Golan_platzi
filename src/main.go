package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPc *pc) duplicate_Ram() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(*b)
	*b = 100

	myPc := pc{ram: 15, disk: 200, brand: "MSI"}

	fmt.Println(myPc)

	myPc.ping()
	myPc.duplicate_Ram()
	fmt.Println(myPc)
}
