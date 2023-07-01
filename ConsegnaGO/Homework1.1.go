/*  Giacomo Calabria
	Homework 1 versione "generica"

	Questo programma conta il numero di volte in cui ciascun carattere
   	compare in una stringa. Il programma utilizzando la concorrenza, avvia una
   	goroutine per ogni carattere della stringa. 
   
   	Utilizza il meccanismo di sincronizzazione delle WaitGroup e un channel per
   	tenere traccia del conteggio totale dei caratteri corrispondenti.

	Rispetto alla versione data nella consegna, utilizza una struttura dati di tipo mappa 
	per tenere traccia del conteggio dei caratteri.
*/

package main

import (
	"fmt"  // Standard package per la stampa su console 
	"sync" // Package sync per la sincronizzazione con WaitGroup
)

func contaCarattere(key rune, counter chan map[rune] int, wg *sync.WaitGroup){
	// Incremento il conteggio utilizzando il Channel 
	mappa := <- counter 	// Prendo la mappa dal channel
	mappa[key]++ 		// Incremento il valore corrispondente al carattere corrente nella mappa. 
						// Se il carattere non è presente, viene aggiunto con valore 0 e poi comunque viene incrementato.
	counter <- mappa   	// Rimetto la mappa aggiornata nel channel
	defer wg.Done() // Decremento il WaitGroup alla fine della goroutine
}

func main(){
	stringa := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed" // Stringa da analizzare

	// **** STRUMENTI PER LA SINCRONIZZAZIONE ****
	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	conta := make(chan map[rune] int, 1) // Channel per il conteggio dei caratteri
	conta <- map[rune] int{} // Inizializzazione del canale con una mappa vuota
	
	// **** ITERAZIONE CARATTERI DELLA STRINGA ****
	for _, elem := range stringa {
        // Incremento il WaitGroup per ogni goroutine creata
		wg.Add(1)
        // Avvio di una goroutine per contare il carattere corrente
		go contaCarattere(elem, conta, &wg) // Passo alla goroutine gli strumenti per la sincronizzazione
	}

	// **** ATTENDO TERMINAZIONE GOROUTINE  ****
	wg.Wait() // Attendo la terminazione di tutte le goroutine
	close(conta) // Chiudo il channel per evitare ev. deadlock 

	// **** STAMPA DEL RISULTATO ****
	totalMap := <- conta // Prendo dal channel la mappa con il conteggio finale dei caratteri
	for key, value := range totalMap {
		fmt.Printf("Il carattere '%c' compare: %d volte nella stringa \"%s\"\n", key, value, stringa) // Stampo il risultato
	}
}

/* OUTPUT:

Il carattere 'm' compare: 3 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'a' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'c' compare: 3 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'o' compare: 4 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'p' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'l' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere ',' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere ' ' compare: 8 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'r' compare: 3 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'e' compare: 6 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 's' compare: 5 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 't' compare: 5 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'n' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'g' compare: 1 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'L' compare: 1 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'i' compare: 6 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'u' compare: 2 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"
Il carattere 'd' compare: 3 volte nella stringa "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed"

*/