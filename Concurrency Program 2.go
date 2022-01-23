/*En este problema no podemos definir una variable global y que una Goroutine la incremente y la otra la muestre sin
que haya una sincronización entre las mismas, podría una de las funciones imprimir el mismo valor varias veces o la que lo incrementa
hacerlo varias veces antes que la otra lo muestre.
Estas dos funciones deben sincronizarse para que cuando la que incrementa el contador lo haga le avise a la que imprime que lo muestre.
Si ejecutamos este programa veremos los números a partir de 0 de 1 en 1 hasta que el operador presione la tecla "enter" y finalice
el programa.*/

package main

import "fmt"

func contar(canal chan int) {
	x := 0
	for {
		canal <- x
		x++
	}
}
func imprimir(canal chan int) {
	var valor int
	for {
		valor = <-canal
		fmt.Println(valor)
	}
}
func main() {
	canal := make(chan int)
	go contar(canal)
	go imprimir(canal)
	var fin string
	fmt.Scanln(&fin)
}
