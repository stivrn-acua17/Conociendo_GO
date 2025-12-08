package controlflujo

import "fmt"

func Defer() { // La ejecución de la linea con defer va al final
	example_defer_1()
	example_defer_2()
}

func example_defer_1() {
	fmt.Println("EJEMPLO DEFER 1")
	defer fmt.Println("Esto esta impreso al final")
	fmt.Println("Esto esta impreso primero")
}

func example_defer_2() {
	fmt.Println("EJEMPLO DEFER 2")
	for i := 0; i < 3; i++ {  
		defer fmt.Println(i)	// Defer con funcionamiento similar a pilas
	}

	fmt.Println("Bucle terminado")
}
