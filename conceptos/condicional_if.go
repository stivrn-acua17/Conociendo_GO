package conceptos
import (
	"fmt"
	"errors"
)

// Condicionales --------------------------------------------------------------------
func Condicional_if() {
	// FUNCIONES ------------------------------------------------------------------------
	decirHola()
	suma := add(2, 3)
	fmt.Println(suma)

		// IF / ELSE -----------------------------------------------------------
	division, err := dividir(1, 1)
	if err != nil { fmt.Println("Hubo un error") }
	fmt.Println(division)

	if y := 10; y > 5 {  // DECLARAR EN EL IF SE LLAMA ALCANCE LIMITADO, UTIL PARA VARIABLES TEMPORALES
		fmt.Println("y es mayor de 5")
	} 
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