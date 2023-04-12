/* Questo programma conta il numero di volte in cui un determinato carattere
   compare in una stringa. Il programma utilizza la concorrenza, avviando una
   goroutine per ogni carattere della stringa. 
   
   Utilizza il meccanismo di sincronizzazione delle WaitGrouo e un channel per
   tenere traccia del conteggio totale dei caratteri corrispondenti.

*/

package main

import (
	"fmt"
	"sync"
)

func increaseNumber(conta chan int) {
	num := <-conta
	num++
	conta <- num
}

func main() {
	// Input string
	stringa := "ccccccccc"

	// Character to find
	carattere := 'c'

	// Creazione degli strumenti per la sincronizzazione / concorrenza

	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	conta := make(chan int, 1) // Channel per il conteggio dei caratteri
	conta <- 0 // Inizializzazione del canale a 0

	// Iterazione sui caratteri della stringa
	for _, elem := range stringa {
		
        // Incremento il WaitGroup per ogni goroutine creata
		wg.Add(1)

        // Avvio di una goroutine per contare il carattere corrente
		go func(goElem rune){
			defer wg.Done()

			// Se il carattere corrisponde al carattere cercato
			if goElem == carattere {
				// Incremento il conteggio utilizzando il Channel
				increaseNumber(conta)
			}
		}(elem)
	}

    // Attendo la terminazione di tutte le goroutine
    wg.Wait()

    // Il Channel conta contiene alla fine il conteggio finale del cartttere
	total := <- conta

	fmt.Printf("Il carattere '%c' compare %d volte nella stringa %s\n", carattere, total, stringa)
}