package estructuradedatos

import "fmt"

// Trabajo con referencias de memoria
// Mas simples y seguros que en otros lenguajes pero igual de seguros
func Pointer1() {
	example_pointer_1()
	example_pointer_2()
	example_pointer_3()
}

func example_pointer_1() {
	fmt.Println("EJEMPLO POINTER 1")
	x := 10
	p := &x

	fmt.Println(p)  // Salida: Direccion de memoria de x
	fmt.Println(*p) // Salida: 10
}

func example_pointer_2() {
	fmt.Println("EJEMPLO POINTER 2")
	x := 20
	p := &x
	*p = 50 // Cambio de valor en la referencia de memoria p

	fmt.Println(x) // Salida: 50
}

func incrementarPorValor(x int) {
	x = x + 1
	fmt.Println("Dentro de la funcion:", x)
}

func incrementarPorPointer(x *int) {
	*x = *x + 1
	fmt.Println("Dentro de la funcion:", *x)
}

func example_pointer_3() {
	fmt.Println("EJEMPLO POINTER 3")
	num := 10

	incrementarPorValor(num)
	fmt.Println("Despues de incrementarPorValor:", num)

	incrementarPorPointer(&num)
	fmt.Println("Despues de incrementarPorPointer:", num)

}
