package main

import (
	"math/rand"
	"fmt"
	"time"
)

func Calcular(valores chan int)  {
	valor := int(rand.Int31n(10))
	time.Sleep(time.Millisecond * 50)
	fmt.Println("Valor aleatorio: ", valor)
	valores <- valor
}

func main()  {
	fmt.Println("Probando canales")
	valores := make(chan int, 3)
	defer close(valores)

	go Calcular(valores)
	go Calcular(valores)
	go Calcular(valores)

	valor := <- valores

	time.Sleep(time.Millisecond * 1000)
	fmt.Println(valor)
}