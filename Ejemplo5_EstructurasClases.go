package main

import "fmt"

func main() {
	p := Persona{"Sebastian", 178}
	fmt.Println(p.correr());
}

type Persona struct {
	nombre string
	estatura float64
}

func (p *Persona) correr()(string, string){
	return p.nombre, "corriendo"
}