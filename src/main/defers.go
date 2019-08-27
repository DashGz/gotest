package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Inicio...")
	defer funcion1()
	funcion2()
	fmt.Println("Fin...")
}

func funcion1()  {
	time.Sleep(time.Second * 5)
	fmt.Println("Proceso 1...")
}

func funcion2()  {
	time.Sleep(time.Second * 10)
	fmt.Println("Proceso 2...")

}


