package controlflujo

import "fmt"

func Continue() {
	example_continue_1()
	example_continue_2()
}

func example_continue_1() {
	fmt.Println("EJEMPLO CONTINUE 1")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue  // Sale de esa iteraccion pero continua el bucle
		}

		fmt.Println(i)  // 1, 3
	}
}

func example_continue_2() {
	fmt.Println("EJEMPLO CONTINUE 2")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue
			}

			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}