package construcciones

import (
	"fmt"
	"time"
)

func Select() { // Especie de switch para canales
	example_select_1()
	example_select_2()
}

func example_select_1() {
	fmt.Println("EJEMPLO SELECT 1")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "Canal 1" }()
	go func() { ch2 <- "Canal 2" }()

	select {
	case msg1 := <- ch1:
		fmt.Println(msg1)
	case msg2 := <- ch2:
		fmt.Println(msg2)
	}
}

func example_select_2() {
	fmt.Println("EJEMPLO SELECT 2")
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hecho"
	}()

	select {
	case msg := <- ch:
		fmt.Println(msg)
	case <- time.After(1 * time.Second): // La routine toma 2 segundos
		fmt.Println("Timeout!") 
		// Ejecucion que se toma demasiado tiempo, por lo que se para y se da espacio
		// a otras peticiones de avanzar y no consumir mas recursos innecesariamente
	}
}