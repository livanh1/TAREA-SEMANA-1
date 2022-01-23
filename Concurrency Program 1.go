/*Vamos por partes, lo primero que hemos hecho ha sido inicializar nuestro WaitGroup, como cualquier otro struct,
después gracias al método Add le decimos cuantas rutinas vamos a ejecutar, o más bien cuantas veces llamaremos al Done()
antes de salir, finalmente gracias al Wait() bloquearemos la ejecución hasta que recibamos los Done() que habíamos indicado anteriormente.
Hay que además hacer un apunte, si nuestro WaitGroup se pasa a una función o método éste deberá ser pasado por referencia,
de lo contrario, estaríamos tratando sobre una copia y nunca recibiríamos el Done() pertinente.*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("I'm a gorutine")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Hello, playground")
}
