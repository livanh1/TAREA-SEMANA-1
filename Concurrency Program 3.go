/*Cuando comienza a ejecutarse el programa desde la main se llama la goroutine mostrar0
Pero la main no se bloquea hasta que finalice la función "mostrar0" sino que continúa con
la ejecución se las siguientes instrucciones, es decir llama a la función "mostrar1":*/
package main

import "fmt"

func mostrar0() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("0-")
	}
}
func mostrar1() {
	for f := 1; f <= 1000; f++ {
		fmt.Print("1-")
	}
}
func main() {
	go mostrar0()
	go mostrar1()
	var continua string
	fmt.Scan(&continua)
}
