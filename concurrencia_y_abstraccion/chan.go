package concurrenciayabstraccion

import "fmt"

// No te comuniques compartiendo memoria, comparte la memoria comunicándote
func Chan() { // Los canales permiten que las goroutines envian y reciban información de manera segura
	example_chan_1()
	example_chan_2()
}

func example_chan_1() {
	fmt.Println("EJEMPLO CHAN 1")
	ch := make(chan int)

	go func() { ch <- 42 }() // Funcion anonima

	val := <-ch
	fmt.Println(val)
}

func example_chan_2() {
	fmt.Println("EJEMPLO CHAN 2")
	ch := make(chan string, 2)
	ch <- "Hola"
	ch <- "Mundo"
	fmt.Println(<-ch) // Salida: Hola
	fmt.Println(<-ch)	// Salida: Mundo
}