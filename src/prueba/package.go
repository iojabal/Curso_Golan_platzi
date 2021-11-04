package prueba

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrintMessage
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessace1(text string) {
	fmt.Println(text)
}
