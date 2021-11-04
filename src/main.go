package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("hello")

	wg.Add(1)
	go say("world", &wg)

	wg.Wait()
	//opcion no recomendable

	go func(text string) {
		fmt.Println(text)
	}("adios")

	time.Sleep(time.Second * 1)

}
