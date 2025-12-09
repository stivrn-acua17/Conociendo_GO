package concurrenciayabstraccion

import (
	"fmt"
	"time"
)

func Go_Routine() { // Concurrencia
	example_go_1()
	example_go_2()
}

func saluda() {
	fmt.Println("Hola desde un goroutine")
}

func example_go_1() {
	fmt.Println("EJEMPLO GO 1")
	go saluda()
	time.Sleep(time.Second) // Tiempo necesario para que la linea anterior se ejecute
	fmt.Println("main function completed")
}

func printNumeros(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100* time.Millisecond)
	}
}

func example_go_2() {
	fmt.Println("EJEMPLO GO 2")
	go printNumeros(5)
	go printNumeros(5)
	time.Sleep(time.Second) // Tiempo necesario para que ambas subrutinas terminen
}