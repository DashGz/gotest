package main

import (
	"fmt"
	"../domain"
)

func main()  {
	articulo := domain.Articulo{

	}

	articulo.SetId()
	fmt.Println(articulo.Id)
}
