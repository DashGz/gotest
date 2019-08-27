package main

import (
	"fmt"
	"time"
)

var counter = 0

func main()  {
	fmt.Println("Inicio...", counter)
	for i := 0; i < 5; i++ {
		go proceso()

	}

	time.Sleep(time.Millisecond * 100)

	fmt.Println("Final...")

}

func proceso()  {
	counter++
	fmt.Println("Procesando...", counter)
}
