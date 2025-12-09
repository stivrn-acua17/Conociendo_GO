package basicos

import "fmt"

func CondicionalElse() {
	example_else_1()
}

func example_else_1() {
	fmt.Println("EJEMPLO ELSE 1")
	x := -1
	
	if x > 0 {
		fmt.Println("es positivo")
	} else if x == 0 {
		fmt.Println("es cero")	
	} else {
		fmt.Println("es negativo")
	}
}