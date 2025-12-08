package basicos

import "fmt"

// En Golang solo hay una manera de hacer bucles y es con for
func Bucle() {
	example_bucle_1()
	example_bucle_2()
}

func example_bucle_1() {
	fmt.Println("EJEMPLO BUCLE 1")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func example_bucle_2() {
	fmt.Println("EJEMPLO BUCLE 2")
	nums := []int{10, 20, 30}
	for _, num := range nums {	// "_" para ignorar el index
		fmt.Println(num)
	}
}