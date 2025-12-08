package construcciones

import "fmt"

// Salta a la etiqueta que se asigne
// Raramente usado porque puede hacer el codigo dificil de leer y mantener
// Casi siempre hay una mejor alternativa que goto
// Se usa en casos especiales de optimización
func Goto() {
	example_goto_1()
	example_goto_2()
}

func example_goto_1() {
	fmt.Println("EJEMPLO GOTO 1")
	fmt.Println("Antes de goto")
	goto End
	fmt.Println("Esto no se ejecutara")
	End:
		fmt.Println("Despues de goto")
}

func example_goto_2() {
	fmt.Println("EJEMPLO GOTO 2")
	i := 0
	Loop: 
		if i < 3 {
			fmt.Println(i)
			i++
			goto Loop
		}
}