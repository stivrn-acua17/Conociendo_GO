package controlflujo

import "fmt"

func Break() {
	example_break_1()
}

func example_break_1() {
	fmt.Println("EJEMPLO BREAK 1")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // Rompe el bucle
		}

		fmt.Println(i)  // 0, 1, 2
	}
}