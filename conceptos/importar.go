package conceptos
import (
	"fmt"
	"aprende-golang/mathutil"
	"math"
)

// Importar -------------------------------------------------------------------------
func Importar() {
	fmt.Println(mathutil.Add(2, 5))
	// fmt.Println(mathutil.resta(3,2)) // No funciona al ser una funcion privada
	fmt.Println(math.Sqrt(144))
}