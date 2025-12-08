package controlflujo

import "fmt"

type MiInt int  // Usarse fuera de la funcion
type Celsius float64

func Type() {
	example_type_1()
	example_type_2()
}

func example_type_1() {
	fmt.Println("EJEMPLO TYPE 1")
	var m MiInt = 5
	fmt.Println(m)
}

func (c Celsius) ToFahrenheit() float64 {	// Funcion declarada para el tipo Celsius
	return float64(c * 9/5 * 32)
}

func example_type_2() {
	fmt.Println("EJEMPLO TYPE 2")
	temp := Celsius(25)
	fmt.Println(temp.ToFahrenheit())
}