package conceptos
import (
	"fmt"
	"aprende-golang/mathutil"
)

// Variables ------------------------------------------------------------------------
func Variables() {
	var x int = 5
	fmt.Println(x)

	var a, b = 3, 3
	fmt.Println(a, b)

	//DECLARACION DE VARIABLES CON :=, SOLO SE PUEDE USAR DENTRO DE UNA FUNCION -------
	y := 5
	fmt.Println(y)

	nombre, edad := "Juan", 32
	fmt.Println(nombre)
	fmt.Println(edad)

	fmt.Println(mathutil.Edad)
}
