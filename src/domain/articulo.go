package domain

import "math/rand"

type Articulo struct {
	Id int
	Name string
}

func NewArticulo(id int, name string)  *Articulo{
	return &Articulo{
		Id: id,
		Name: name,
	}
}

func (a *Articulo) SetId()  {
	a.Id = int(rand.Int())
}