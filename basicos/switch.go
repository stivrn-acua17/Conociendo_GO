package basicos

import "fmt"

func Switch() {
	example_switch_1()
	example_switch_2()
	example_switch_3()
}

func example_switch_1() {
	fmt.Println("EJEMPLO SWITCH 1")
	dia := 3
	
	switch dia {
	case 1: 
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	default:
		fmt.Println("Otro dia")
	}
}

func example_switch_2() {
	fmt.Println("EJEMPLO SWITCH 2")
	letter := "a"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vocal")
	default:
		fmt.Println("Consonante")
	} 
}

func example_switch_3() {
	fmt.Println("EJEMPLO SWITCH 3")
	numero := 0

	switch numero {
	case 0, 1:
		fmt.Println("Numero es 0 o 1")
	default:
		fmt.Println("Numero es cualquier otro")
	}
}