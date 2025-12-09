package estructuradedatos

import "fmt"

// En golang, los arrays son colecciones de elementos del mismo tipo con un tamaño fijo
func Array() {
	example_array_1()
	example_array_2()
}

func example_array_1() {
	fmt.Println("EJEMPLO ARRAY 1")
	
	var arr [5]int
	arr[0] = 10
	arr[1] = 20

	fmt.Println(arr)
	fmt.Println(arr[1])
}

func example_array_2() {
	fmt.Println("EJEMPLO ARRAY 2")
	arr := [3]string{"Go", "Rust", "Python"}

	for i, v := range arr {
		fmt.Printf("Index %d: %s\n", i, v)
	}
}