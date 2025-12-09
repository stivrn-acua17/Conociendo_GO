package estructuradedatos

import "fmt"

// Trabajo con referencias de memoria
// Mas simples y seguros que en otros lenguajes pero igual de seguros
func Pointer3() {
	example_pointer3_1()
}

type CuentaDeBanco struct {
	Dueño string
	Balance float64
}

// Método con receptor por valor - NO modifiica el original
func (ba CuentaDeBanco) DepositoPorValor(amount float64) {
	ba.Balance += amount
	fmt.Printf("Saldo después del depósito (metodo por valor): %.2f\n", ba.Balance)
}

// Método con receptor por pointer - SI modifica el original
func (ba *CuentaDeBanco) Deposito(cantidad float64) {
	ba.Balance += cantidad
	fmt.Printf("Saldo después del depósito: %.2f\n", ba.Balance)
}

func (ba *CuentaDeBanco) Retirar(cantidad float64) {
	if ba.Balance >= cantidad {
		ba.Balance -= cantidad
		fmt.Printf("Retiro exitoso. Saldo actual: %.2f\n", ba.Balance)
	} else {
		fmt.Println("Fondos insuficientes")
	}
}

func example_pointer3_1() {
	account := CuentaDeBanco {
		Dueño: "Juan Pérez",
		Balance: 1000.0,
	}

	fmt.Printf("Saldo inicial: %.2f\n", account.Balance) // 1000

	// Método por valor - no modifica el original
	account.DepositoPorValor(100.0)
	fmt.Printf("Saldo real después de DepositoPorValor: %.2f\n", account.Balance) // 1000

	// Métodos por pointer - modifican el original
	account.Deposito(200.0)
	fmt.Printf("Saldo real después de Deposito: %.2f\n", account.Balance) // 1200.0

	account.Retirar(150.0)
	fmt.Printf("Saldo final: %.2f\n", account.Balance) // 1050.0
}