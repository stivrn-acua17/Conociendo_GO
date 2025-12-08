package conceptos

import "fmt"

func Condicional_else() {
	x := -1
	
	if x > 0 {
		fmt.Println("es positivo")
	} else if x == 0 {
		fmt.Println("es cero")	
	} else {
		fmt.Println("es negativo")
	}
}