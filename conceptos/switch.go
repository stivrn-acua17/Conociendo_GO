package conceptos

import "fmt"

func Switch_func() {
	dia := 3
	
	switch dia {
	case 1: 
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Otro dia")
	}
}