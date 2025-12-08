package conceptos

import "fmt"

// EN GO SOLO HAY UNA MANERA DE HACER UN BUCLE Y ES UN FOR
func Bucle() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	nums := []int{10, 20, 30}
	for _, num := range nums {	// "_" para ignorar el index
		fmt.Println(num)
	}
}