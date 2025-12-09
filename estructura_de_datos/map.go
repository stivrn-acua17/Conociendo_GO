package estructuradedatos

import "fmt"

// En golang, es similar a un diccionario con su clave, valor
// En golang, map no garantizan orden, los mapas son referencias
func Map() {
	example_map_1()
	example_map_2()
}

func example_map_1() { 
	fmt.Println("EJEMPLO MAP 1")
	m := map[string]int{"manzana": 2, "banana": 3}
	fmt.Println(m["manzana"])

	for key, value := range m {
		fmt.Printf("key es: %s; value es: %d\n", key, value)
	}
}

func example_map_2() {
	fmt.Println("EJEMPLO MAP 2")
	m := make(map[string]int)
	
	m["manzana"] = 2
	m["banana"] = 3
	fmt.Println(m["manzana"])
}