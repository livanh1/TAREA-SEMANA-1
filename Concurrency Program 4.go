package main

import "fmt"

func mostrarImpar() {
	for f := 1001; f <= 2000; f++ {

		if f%2 != 0 {

			fmt.Println(f, " es un numero impar!")

		}

	}
}
func mostrarPar() {
	for f := 1; f <= 1000; f++ {

		if f%2 == 0 {

			fmt.Println(f, " es un numero par!")

		}

	}
}
func main() {
	go mostrarImpar()
	go mostrarPar()
	var continua string
	fmt.Scan(&continua)
}
