package estructuradedatos

import "fmt"

// Trabajo con referencias de memoria
// Mas simples y seguros que en otros lenguajes pero igual de seguros
func Pointer2() {
	example_pointer2_1()
}

type LargeStruct struct {
	Data [1000]int
	Name string
	ID int
}

// Sin pointer - copia de toda la estructura (ineficiente)
func processByValue(ls LargeStruct) {
	fmt.Println("Procesando:", ls.Name)
	// Aqui se copiaron 1000 enteros + string + int
}

// Con pointer - solo pasa la dirección (eficiente)
func processByPointer(ls *LargeStruct) {
	fmt.Println("Procesando:", ls.Name)
	// Solo se paso una direccion de memoria
}

func example_pointer2_1() {
	large := LargeStruct {
		Name: "Mi estructura grande",
		ID: 1,
	}

	processByValue(large)
	processByPointer(&large)
}