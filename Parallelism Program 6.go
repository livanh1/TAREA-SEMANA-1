/*Imagina que tienes que escribir unos cuantos correos electrónicos. Van a ser largos
y laboriosos, y la mejor forma de entretenerse es escuchar música mientras se escriben,
es decir, escuchar música “en paralelo” a la redacción de los emails.*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

// Tarea que se hará con el tiempo.
func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1.")
	wg.Done()
}
func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2.")
	wg.Done()
}
func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3.")
	wg.Done()
}

// Tarea realizada en paralelo
func listenForever() {
	for {
		printTime("Listening...")
	}
}
func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go listenForever()

	// Darle algo de tiempo para que listenForever comience
	time.Sleep(time.Nanosecond * 10)

	// Escribiendo los emails
	go writeMail1(&waitGroup)
	go writeMail2(&waitGroup)
	go writeMail3(&waitGroup)

	waitGroup.Wait()
}
