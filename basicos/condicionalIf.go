package basicos

import (
	"fmt"
	"errors"
)

// Condicionales --------------------------------------------------------------------
func CondicionalIf() {
	// Funciones ----------------------------------------------------------------------
	decirHola()
	suma := add(2, 3)
	fmt.Println(suma)

	example_if_1()
	example_if_2()
}

func decirHola() {
	fmt.Println("Hola!")
}

func add(x, y int) int {
	return x + y
}

func dividir(x, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("no puede ser cero")
	}

	return x / y, nil
}

func example_if_1() {
	fmt.Println("EJEMPLO IF 1")
	if y := 10; y > 5 {  // Declarar en el if una variable se denomida alcance limitado, util para variables temporales
		fmt.Println("y es mayor de 5")
	} 
}

func example_if_2() {
	fmt.Println("EJEMPLO IF 2")
	division, err := dividir(1, 1)
	if err != nil { fmt.Println("Hubo un error") }
	fmt.Println(division)
}