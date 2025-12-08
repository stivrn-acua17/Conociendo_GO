package controlflujo

import "fmt"

func Range() {
	example_range_1()
	example_range_2()
}

func example_range_1() {
	fmt.Println("EJEMPLO RANGE 1")
	nums := []int{1, 2, 3, 4}

	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}

func example_range_2() {
	fmt.Println("EJEMPLO RANGE 2")
	m := map[string]int{"a": 1, "b": 2}

	for clave, valor := range m {
		fmt.Printf("Key: %s, Value: %d\n", clave, valor)
	}
}