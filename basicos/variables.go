package basicos

import (
	"fmt"
	"aprende-golang/mathutil"
)

// Variables ------------------------------------------------------------------------
func Variables() {
	example_variables_1()
	example_variables_2()	
	example_variables_3()
	example_variables_4()
	example_variables_5()
}

func example_variables_1() {
	fmt.Println("EJEMPLO VARIABLES 1")
	var x int = 5
	fmt.Println(x)
}

func example_variables_2() {
	fmt.Println("EJEMPLO VARIABLES 2")
	var a, b = 3, 3
	fmt.Println(a, b)
}

func example_variables_3()  {
	fmt.Println("EJEMPLO VARIABLES 3")

	// Declaración de variables con :=, solo se puede usar dentro de una función
	y := 5
	fmt.Println(y)
}

func example_variables_4() {
	fmt.Println("EJEMPLO VARIABLES 4")
	nombre, edad := "Juan", 32
	fmt.Println(nombre)
	fmt.Println(edad)
}

func example_variables_5() {
	fmt.Println("EJEMPLO VARIABLES 5")
	fmt.Println(mathutil.Edad) // Variable llamada desde otro archivo
}