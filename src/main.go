package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPV pc) String() string {
	return fmt.Sprintf("tengo %d GB Ram, %d de dsico y es %s", myPV.ram, myPV.disk, myPV.brand)
}

func main() {
	myPc := pc{ram: 15, disk: 200, brand: "MSI"}

	fmt.Println(myPc)

}
