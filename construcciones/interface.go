package construcciones

import "fmt"

type Hablador interface {
	Hablar() string
}

type Perro struct {}

func (p Perro) Hablar() string {
	return "Woof!!"
}

func Interface() {
	example_interface_1()
	example_interface_2()
}

// Con interface podriamos aplicar polimorfismo si un tipo tiene las mismas funciones
func example_interface_1() {
	fmt.Println("EJEMPLO INTERFACE 1")
	var p Hablador = Perro {}

	fmt.Println(p.Hablar())
}

// Al no tener ningun metodo, todo lo satisface, su uso no se recomienda, usar de manera moderada
func imprimirCualquierCosa(v interface {}) {
	fmt.Println(v)
}

func example_interface_2() {
	fmt.Println("EJEMPLO INTERFACE 2")
	imprimirCualquierCosa(42)
	imprimirCualquierCosa("hola")
	imprimirCualquierCosa(3.14)
}