package main

import (
	"os"
	"strconv"
	"fmt"
	"errors"
)

func main()  {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := convertir(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	fmt.Println(n)
}

func convertir(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1, errors.New("El argumento es un número no válido")
	}
	return n, nil
}
