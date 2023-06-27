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

func contaCarattere(current rune, test rune, conta chan int, wg *sync.WaitGroup){
	// Se il carattere corrisponde al carattere cercato
	if current == test {
		// Incremento il conteggio utilizzando il Channel
		increaseNumber(conta)
	}

	defer wg.Done() // Decremento il WaitGroup alla fine della goroutine
}

func main() {
	stringa := "ciao come stai?" // Stringa da analizzare

	carattere := 'a' // Carattere da cercare

	// Creazione degli strumenti per la sincronizzazione
	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	conta := make(chan int, 1) // Channel per il conteggio dei caratteri
	conta <- 0 // Inizializzazione del canale a 0

	// Iterazione sui caratteri della stringa
	for _, elem := range stringa {	
        // Incremento il WaitGroup per ogni goroutine creata
		wg.Add(1)
        // Avvio di una goroutine per contare il carattere corrente
		go contaCarattere(carattere, elem, conta, &wg)
	}

    // Creazione di una goroutine per attendere la terminazione di tutte le goroutine
	go func(){
		wg.Wait() // Attendo la terminazione di tutte le goroutine
		close(conta) // Chiudo il channel
	}()

    // Il channel conta contiene alla fine il conteggio finale del cartttere
	total := <- conta // Prendo il valore dal channel

	fmt.Printf("Il carattere '%c' compare %d volte nella stringa %s\n", carattere, total, stringa)
}
