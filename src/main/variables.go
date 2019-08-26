package main

import (
	"fmt"
)

func main() {
	var paginas int = 300
	autor := "Pepe"
	autor = "Pipo"
	autor, pagina := "Jhon", 10
	color := getColor()

	c, _ := sumar(10, 20)

	fmt.Println(paginas, autor, pagina, color, c)
}

func getColor() int {
	return 123
}

func getNada()  {

}

func sumar(a, b int) (c int, d bool) {
	/*c := a + b
	d := true

	return c, d*/

	c = a + b
	d = true
	return
}
