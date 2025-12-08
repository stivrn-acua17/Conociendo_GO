package conceptos
import "fmt"

const Pi = 3.14 // GLOBAL
const (
	A = 1
	B = 2
)

// Constantes -----------------------------------------------------------------------
func Constantes() {
	// LLAMANDO CONSTANTE GLOBAL ------------------------------------------------------
	fmt.Println(Pi)
	fmt.Println(A + B)

}