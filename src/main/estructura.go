package main

import "fmt"

type Capitulo struct {
	Titulo string
	Paginas int
}

func main()  {
	cap1 := Capitulo{
		Titulo: "Introducción",
		Paginas: 20,
	}

	fmt.Println(cap1.Titulo, cap1.Paginas)
}

