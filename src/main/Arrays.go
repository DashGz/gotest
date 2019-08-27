package main

import (
	"math/rand"
	"sort"
	"fmt"
)

func main()  {

	//make es como un new y ademas lo inicializa
	//puntaje:= make([]int, 10)

	//si a la linea de arriba le agregamos un 0, creamos un array de largo 10 con un slice
	//pero no inicia el espacio en memoria

	//puntaje:= make([]int, 0, 10)

	//esta linea da error porque no inicializamos nunca el espacio en memoria
	// puntaje[5] = 1000

	//para agregar un elemento inicializado usamos append
	/*
	puntaje = append(puntaje, 5)
	fmt.Println(puntaje)

	puntaje = puntaje[0:6]
	puntaje[5] = 1000
	fmt.Println(puntaje)



	puntaje:= make([] int, 0, 5)
	l := len(puntaje)
	c := cap(puntaje)
	fmt.Println(l, c)

	for i := 0; i<25; i++{
		puntaje = append(puntaje, i)
		if cap(puntaje) != c {
			l = len(puntaje)
			c = cap(puntaje)
			fmt.Println(l, c)
		}
	}
	*/

	p:= make([] int, 100)

	for i := 0; i<100; i++ {
		p[i] = int(rand.Int())
	}

	sort.Ints(p)

	copia := make([] int, 10)
	copy(copia, p[20:30])

	fmt.Println(copia)
}
