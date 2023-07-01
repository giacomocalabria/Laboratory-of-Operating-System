/*	Author: Giacomo Calabria
	Homework 2

	Questo programma simula un'agenzia di noleggi d'auto, gestisce le prenotazioni di 10 clienti.
	Un cliente nolegga un veicolo dai disponibili: Berlina, SUV o Station Wagon.
	Ogni cliente può noleggiare un veicolo contemporaneamente ad altri.

	Utilizza il meccanismo di sincronizzazione delle WaitGroup e un channel di tipo map[int] int per tenere traccia del conteggio totale dei veicoli noleggiati per tipo.
*/

package main

import (
	"fmt"	// Standard package per la stampa su console 
	"time"	// Package time per inizializzare il generatore di numeri casuali
	"math/rand" // Package rand per la generazione di numeri casuali
	"sync"  // Package sync per la sincronizzazione con WaitGroup
)

// Struttura Cliente con un campo "nome" di tipo string
type Cliente struct{
	nome string
}

// Struttura Veicolo con un campo "tipo" di tipo string
type Veicolo struct{
	tipo string
}

// Funzione ausiliaria per aumentare il valore di un canale di tipo map[int] int
func increaseNumber(counter chan map[int] int, key int) {
	// Incremento il conteggio utilizzando il Channel 
	mappa := <- counter 	// Prendo la mappa dal channel
	mappa[key]++ 		// Incremento il valore corrispondente al carattere corrente nella mappa. 
						// Se il carattere non è presente, viene aggiunto con valore 0 e poi comunque viene incrementato.
	counter <- mappa   	// Rimetto la mappa aggiornata nel channel
}

//La  function  "noleggia"  che  prende  come  input  un cliente  e  che  prenota uno a caso tra i veicoli. Questa function deve anche stampare che il cliente x ha noleggiato il veicolo y.
func noleggia(c Cliente, veicoli []Veicolo, conta chan map[int] int, wg *sync.WaitGroup) {
	i := rand.Intn(len(veicoli))
	increaseNumber(conta, i)
	fmt.Printf("Il Cliente %s ha noleggiato il veicolo %s.\n", c.nome, veicoli[i].tipo)

	defer wg.Done()
}

//La  function  "stampa"  che,  alla  fine  del  processo,  stampa  il  numero  di Berline, SUV e Station Wagon noleggiati.
func stampa(conta chan map[int] int){
	totalMap := <- conta
	fmt.Println("Numero di veicoli noleggiati per tipo:")
	fmt.Printf("Berline-> %d\n", totalMap[0])
	fmt.Printf("SUV-> %d\n", totalMap[1])
	fmt.Printf("Station Wagon-> %d\n", totalMap[2])
}

func main(){
	rand.Seed(time.Now().UnixNano()) // Inizializzo il generatore di numeri casuali

	// **** INIZIALIZZAZIONE DATI ****
	clienti := []Cliente{
		{"A"},{"B"},{"C"},{"D"},{"E"},{"F"},{"G"},{"H"},{"I"},{"L"},
	}

	veicoli := []Veicolo{
		{"Berlina"},{"SUV"},{"Station Wagon"},
	}

	// **** STRUMENTI PER LA SINCRONIZZAZIONE ****
	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	conta := make(chan map[int] int, 1) // Creo un Channel per il conteggio dei veicoli noleggiati
	conta <- map[int] int{} // Inizializzazione del canale con una mappa vuota

	// **** ITERAZIONE DEI CLIENTI ****
	for _, cliente := range clienti{
		wg.Add(1) // Aggiungo un'attesa per ogni goroutine
		go noleggia(cliente, veicoli, conta, &wg)
	}

	// **** ATTENDO TERMINAZIONE GOROUTINE  ****
	wg.Wait() // Attendo la terminazione di tutte le goroutine
	close(conta) // Chiudo il channel per evitare ev. deadlock 

	// **** STAMPA RISULTATI ****
	stampa(conta)
}


/* ESECUZIONE

Il Cliente L ha noleggiato il veicolo Berlina.
Il Cliente A ha noleggiato il veicolo Berlina.
Il Cliente F ha noleggiato il veicolo Berlina.
Il Cliente G ha noleggiato il veicolo SUV.
Il Cliente H ha noleggiato il veicolo Station Wagon.
Il Cliente C ha noleggiato il veicolo SUV.
Il Cliente D ha noleggiato il veicolo Station Wagon.
Il Cliente E ha noleggiato il veicolo Berlina.
Il Cliente B ha noleggiato il veicolo Station Wagon.
Il Cliente I ha noleggiato il veicolo Station Wagon.
Numero di veicoli noleggiati per tipo:
Berline-> 4
SUV-> 2
Station Wagon-> 4

*/