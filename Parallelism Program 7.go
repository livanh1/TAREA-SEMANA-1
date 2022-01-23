//Uso de múltiples núcleos simultáneamente. En este caso 2
package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0
}
func main() {
	runtime.GOMAXPROCS(2) //2 núcleos

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
}
