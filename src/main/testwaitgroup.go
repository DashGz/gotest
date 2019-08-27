package main

import (
	"sync"
	"fmt"
	"net/http"
)

/*
import (
	"fmt"
	"sync"
)

func funcion(wg * sync.WaitGroup)  {
	fmt.Println("Dentro de la goroutine")
	wg.Done()
}

func main()  {
	fmt.Println("Inicio...")

	var waitgroup sync.WaitGroup

	for i := 0; i<10 ; i++{
		waitgroup.Add(1)
		go funcion(&waitgroup)
	}

	waitgroup.Wait()
	fmt.Println("Fin de la ejecución")
}
*/

var urls = []string {
	"https://www.google.com.ar",
	"https://www.mercadolibre.com/",
	"https://play.golang.org/",

}

func recuperar(url string, wg *sync.WaitGroup) (string, error)  {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homepage(w http.ResponseWriter, r *http.Request)  {
	var wg sync.WaitGroup
	for _, url := range urls{
		wg.Add(1)
		go recuperar(url, &wg)
	}

	wg.Wait()
	fmt.Println("Respuesta recibida.")
}

func manejador()  {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)
}

func main()  {
	manejador()
}