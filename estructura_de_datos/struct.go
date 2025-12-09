package estructuradedatos

import "fmt"

type Persona struct { // Parecida a las clases en otros lenguajes
	Nombre string
	Edad int
}

type Rectangulo struct {
	Width, Height float64
}

func Struct() {
	example_struct_1()
	example_struct_2()
}

func example_struct_1() {
	fmt.Println("EJEMPLO STRUCT 1")
	p := Persona {
		Nombre: "Cesar",
		Edad: 29,
	}

	nombre := p.Nombre
	fmt.Println(nombre)
	fmt.Println(p)
}

func (r Rectangulo) Area() float64 {
	return r.Width * r.Height
}

func example_struct_2() {
	fmt.Println("EJEMPLO STRUCT 2")
	rect := Rectangulo {
		Width: 5,
		Height: 5,
	}

	fmt.Printf("Area: %f", rect.Area())
}
