package main

import "fmt"


//Estructura

type Persona struct {
	Nombre string
	Edad int
}

type Estudiante struct {
	Persona *Persona
	Legajo int
}

//Constructores

//devuelve un puntero que retorna la dirección de memoria a esa Persona
func newPersona(nombre string, edad int) *Persona {
	return &Persona{
		Nombre:nombre,
		Edad:edad,
	}
}

//Recivers

//asi se declara un metodo para la estructura de arriba, si el nombre esta declarado en minuscula no se puede acceder
//al metodo desde otro archivo, es un modo privado. Los atributos al comenzar con mayuscula si estan disponibles.
func (p *Persona) presentarse()  {
	fmt.Printf("Hola, mi nombre es %s\n", p.Nombre)
}

//este es accesible de otros archivos
func (p *Persona) DecirEdad()  {
	fmt.Printf("Mi edad es %d\n", p.Edad)
}



func main()  {
	john := Persona{
		Nombre: "John",
		Edad: 20,
	}

	prueba(john)
	fmt.Println(john.Edad)

	//el ampersan se usa para indicar que es un puntero
	john2 := &Persona{
		Nombre: "John",
		Edad: 20,
	}

	// por mas que sea un puntero, podemos acceder a los atributos de la misma manera
	prueba2(john2)
	fmt.Println(john2.Edad)

	john2.presentarse()
	john2.DecirEdad()


	//crear un objeto con el constructor

	jane:= newPersona("Jane", 20)
	jane.presentarse()

	//crear un estudiante
	estudiante:= Estudiante{
		Persona:&Persona{
			Nombre: "Pepe",
			Edad: 20,
		},
		Legajo: 1234,
	}

	estudiante.Persona.presentarse()
}


/*Las funciones hacen copias, esta no trabaja sobre el objeto que le enviamos, sólo modifica la copia.
*/
func prueba(p Persona)  {
	p.Edad++
	fmt.Printf("La edad de la copia es %d\n", p.Edad)
}


//el asterisco indica que recibe un puntero
func prueba2(p *Persona)  {
	p.Edad++
	fmt.Printf("La edad de la copia es %d\n", p.Edad)
}