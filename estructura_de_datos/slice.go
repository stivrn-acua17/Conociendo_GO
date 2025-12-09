package estructuradedatos

import "fmt"

func Slice() {
	example_slice_1()
	example_slice_2()
}

func example_slice_1() {
	fmt.Println("EJEMPLO SLICE 1")

	slice := []int{1, 2, 3} // Declarado como array sin especificar su tamaño
	slice = append(slice, 4, 5)
	fmt.Println(slice)
}

func example_slice_2() {
	fmt.Println("EJEMPLO SLICE 2")

	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // Slice con partes de un array, indice final es exclusivo, no incluye el indice 4
	fmt.Println(slice)
}