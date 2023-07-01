/*  Author: Giacomo Calabria
	Homework 1

	Questo programma conta il numero di volte in cui un determinato carattere
   	compare in una stringa. Il programma utilizzando la concorrenza, avvia una
   	goroutine per ogni carattere della stringa. 
   
   	Utilizza il meccanismo di sincronizzazione delle WaitGrouo e un channel per
   	tenere traccia del conteggio totale dei caratteri corrispondenti.
*/

package main

import (
	"fmt"  // Standard package per la stampa su console 
	"sync" // Package sync per la sincronizzazione con WaitGroup
)

func increaseNumber(conta chan int) {
	num := <-conta 	// Prendo il valore dal channel
	num++ 			// Incremento il valore
	conta <- num   	// Rimetto il valore nel channel
}

func contaCarattere(current rune, test rune, conta chan int, wg *sync.WaitGroup){
	if current == test {
		increaseNumber(conta) // Incremento il conteggio utilizzando il Channel
	}

	defer wg.Done() // Decremento il WaitGroup alla fine della goroutine
}

func main(){
	stringa := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed" // Stringa da analizzare

	carattere := 'o' // Carattere da cercare

	// **** STRUMENTI PER LA SINCRONIZZAZIONE ****
	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	conta := make(chan int, 1) // Channel per il conteggio dei caratteri
	conta <- 0 // Inizializzazione del canale a 0 

	// **** ITERAZIONE CARATTERI DELLA STRINGA ****
	for _, elem := range stringa {
        // Incremento il WaitGroup per ogni goroutine creata
		wg.Add(1)
        // Avvio di una goroutine per contare il carattere corrente
		go contaCarattere(carattere, elem, conta, &wg) // Passo alla goroutine gli strumenti per la sincronizzazione
	}

	// **** ATTENDO TERMINAZIONE GOROUTINE  ****
	wg.Wait() // Attendo la terminazione di tutte le goroutine
	close(conta) // Chiudo il channel per evitare ev. deadlock 

	// **** STAMPA DEL RISULTATO ****
	total := <- conta // Prendo dal channel "conta" il conteggio finale del carattere
	fmt.Printf("Il carattere '%c' compare %d volte nella stringa %s\n", carattere, total, stringa) // Stampo il risultato
}


/* OUTPUT

Il carattere 'o' compare 4 volte nella stringa Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed

*/