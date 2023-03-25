package main

import (
	"fmt"
	"time"
)

func main() {
	// canale per comunicare tra i pasticceri
	tortaChan := make(chan int, 5)

	// pasticcere 1: cucina le torte
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Pasticcere 1 - Inizio cottura torta %d\n", i)
			time.Sleep(1 * time.Second) // ci mette 1 secondo a cucinare una torta
			fmt.Printf("Pasticcere 1 - Fine cottura torta %d\n", i)
			tortaChan <- i // mette la torta sul canale
		}
	}()

	// pasticcere 2: guarnisce le torte
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Pasticcere 2 - Inizio guarnizione torta %d\n", i)
			torta := <-tortaChan // prende una torta dal canale
			time.Sleep(4 * time.Second) // ci mette 4 secondi a guarnire una torta
			fmt.Printf("Pasticcere 2 - Fine guarnizione torta %d\n", torta)
			tortaChan <- torta // rimette la torta sul canale
		}
	}()

	// pasticcere 3: decora le torte
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Pasticcere 3 - Inizio decorazione torta %d\n", i)
			torta := <-tortaChan // prende una torta dal canale
			time.Sleep(8 * time.Second) // ci mette 8 secondi a decorare una torta
			fmt.Printf("Pasticcere 3 - Fine decorazione torta %d\n", torta)
		}
	}()

	// attende la fine della produzione delle torte
	time.Sleep(30 * time.Second)
	fmt.Println("Produzione torte completata!")
}