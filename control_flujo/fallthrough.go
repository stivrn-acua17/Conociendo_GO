package controlflujo

import "fmt"

func Fallthrough() { // Permite que un switch corra tanto su linea de ejecucion como la siguiente
	example_fallthrough_1()
	example_fallthrough_2()
}

func example_fallthrough_1() {
	fmt.Println("EJEMPLO FALLTHROUGH 1")
	x := 1

	switch x {
	case 1:
		fmt.Println("Caso 1")
		fallthrough
	case 2:
		fmt.Println("Caso 2") // Caso 1 - Caso 2
	}
}

func example_fallthrough_2() {
	fmt.Println("EJEMPLO FALLTHROUGH 2")
	x := 1

	switch x {
	case 1:
		fmt.Println("Empieza caso 1")
		fallthrough
	case 2:
		fmt.Println("Termina caso 1 y caso 2") // Empieza caso 1 - Termina caso 1 y caso 2
	case 3:
		fmt.Println("Caso 3")
	}
}