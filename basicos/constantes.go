package basicos

import "fmt"

const Pi = 3.14 // GLOBAL
const (
	A = 1
	B = 2
)

// Constantes -----------------------------------------------------------------------
func Constantes() {
	example_constante_1()
}


func example_constante_1() {
	// LLAMANDO CONSTANTE GLOBAL ------------------------------------------------------
	fmt.Println("EJEMPLO IF 1")
	fmt.Println(Pi)
	fmt.Println(A + B)
}