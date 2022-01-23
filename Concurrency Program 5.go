package main

import "fmt"

func mostrarSuma(a int, b int) {
	for f := 1; f <= 10; f++ {

		fmt.Println("La suma de los numeros es: ", ((a + b) * f))

	}
}
func mostrarResta(a int, b int) {
	for f := 1; f <= 10; f++ {

		fmt.Println("La resta de los numeros es: ", ((a - b) * f))

	}
}
func main() {
	go mostrarSuma(5, 6)
	go mostrarResta(1, 2)
	var continua string
	fmt.Scan(&continua)
}
