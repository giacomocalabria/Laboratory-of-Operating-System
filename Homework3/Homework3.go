/*	Author: Giacomo Calabria
	Homework 3

	Questo programma simula la produzione di 5 torte da parte di 3 pasticceri.
	La produzione di ogni torta richiede 3 fasi che devono avvenire in ordine: prima la torta viene cucinata, poi guarnita e infine decorata. I tre pasticceri lavorano contemporaneamente. 

	Utilizza dei Channel per sincronizzare i pasticceri e per tenere traccia delle torte prodotte. Ogni pasticcere è stato implementato come una funzione go che viene eseguita in una goroutine.
*/

package main

import (
	"fmt"  // Standard package per la stampa su console 	
	"time" // Package time per la gestione degli intervalli di tempo
)

/* **** FUNZIONE PRIMO PASTICCERE ****
 * Il pasticcere #1 cucina le torne in un secondo per torta.
 * Ha a disposizione 2 spazi per appoggiare le torte cucinate. Può iniziare la 
 * torta successiva solo se c'è uno spazio libero. Non deve aspettare il pasticcere #2
*/
func pasticcere1(spazi1 chan int, TorteCucinate chan int){
	for i := 1; i <= 5; i++{
		spazi1 <- i  // Occupo uno spazio per appoggiare la torta
		fmt.Printf("Pasticcere 1 - Inizio cottura torta %d\n", i)
		time.Sleep(1 * time.Second)
		fmt.Printf("Pasticcere 1 - Fine cottura torta %d\n", i)
		TorteCucinate <- i // Segnalo che la torta è stata cucinata
	}
	fmt.Println("Cottura torte completata!")
}

/* **** FUNZIONE SECONDO PASTICCERE ****
 * Il pasticcere #2 guarnisce le torte in 4 secondi per torta. 
 * Ha a disposizione 2 spazi per appoggiare le torte guarnite.
*/
func pasticcere2(spazi1 chan int, spazi2 chan bool, TorteCucinate chan int, TorteGuarnite chan int){
	for i := 1; i <= 5; i++{
		torta := <- TorteCucinate // Aspetto che una torta sia cucinata
		spazi2 <- true // Occupo uno spazio per appoggiare la torta
		<- spazi1 // Libero uno spazio per far cucinare la torta successiva
		fmt.Printf("Pasticcere 2 - Inizio guarnizione torta %d\n", torta)
		time.Sleep(4 * time.Second)
		fmt.Printf("Pasticcere 2 - Fine guarnizione torta %d\n", torta)
		TorteGuarnite <- torta // Segnalo che la torta è stata guarnita
	}
	fmt.Println("Guarnizione torte completata!")
}

/* **** FUNZIONE TERZO PASTICCERE ****
 * Il pasticcere #3 si occupa di decorare le torte in 8 secondi per torta.
*/
func pasticcere3(spazi2 chan bool, TorteGuarnite chan int, TorteFinite chan int){
	for i := 1; i <= 5; i++{
		torta := <-TorteGuarnite // Aspetto che una torta sia guarnita
		<- spazi2 // Libero uno spazio per far guarnire la torta successiva
		fmt.Printf("Pasticcere 3 - Inizio decorazione torta %d\n", torta)
		time.Sleep(8 * time.Second)
		fmt.Printf("Pasticcere 3 - Fine decorazione torta %d\n", torta)
		TorteFinite <- torta // Segnalo che la torta è stata decorata ed è finita
	}
}

func main() {
	// **** STRUMENTI PER LA SINCRONIZZAZIONE ****

	// Canale per gli spazi liberi per appoggiare le torne una volta che hanno finito di cucinarle. Se ci sono spazi liberi, può iniziare a cucinare la torta successiva senza aspettare che il secondo pasticcere si liberi per guarnire quella appena cucinata.
	spazi1 := make(chan int, 2)
	// Canale per sapere se c'è una torna da cucinare
	TorteCucinate := make(chan int, 5)

	// Canale per gli spazi libero per appoggiare le torte una volta che ha finito di guarnirle. 
	spazi2 := make(chan bool, 2)

	// Canale per sapere se c'è una torta da farcire
	TorteGuarnite := make(chan int, 5)

	// Canale per sapere se una torta è finita
	TorteFinite := make(chan int, 5)

	// **** INIZIO ESECUZIONE DELLE GOROUTINE****
	fmt.Println("Produzione torte iniziata!")

	go pasticcere1(spazi1, TorteCucinate)
	go pasticcere2(spazi1, spazi2, TorteCucinate, TorteGuarnite)
	go pasticcere3(spazi2, TorteGuarnite, TorteFinite)

	for i := 1; i <= 5; i++{
		torta := <- TorteFinite
		fmt.Printf("Torta %d finita\n", torta)
	}

	fmt.Println("Produzione di tutte le torte finita!")
}



/*  Output di esempio:

Produzione torte iniziata!
Pasticcere 1 - Inizio cottura torta 1
Pasticcere 1 - Fine cottura torta 1
Pasticcere 1 - Inizio cottura torta 2
Pasticcere 2 - Inizio guarnizione torta 1
Pasticcere 1 - Fine cottura torta 2
Pasticcere 1 - Inizio cottura torta 3
Pasticcere 1 - Fine cottura torta 3
Pasticcere 2 - Fine guarnizione torta 1
Pasticcere 2 - Inizio guarnizione torta 2
Pasticcere 3 - Inizio decorazione torta 1
Pasticcere 1 - Inizio cottura torta 4
Pasticcere 1 - Fine cottura torta 4
Pasticcere 2 - Fine guarnizione torta 2
Pasticcere 2 - Inizio guarnizione torta 3
Pasticcere 1 - Inizio cottura torta 5
Pasticcere 1 - Fine cottura torta 5
Cottura torte completata!
Pasticcere 2 - Fine guarnizione torta 3
Pasticcere 3 - Fine decorazione torta 1
Pasticcere 3 - Inizio decorazione torta 2
Pasticcere 2 - Inizio guarnizione torta 4
Torta 1 finita
Pasticcere 2 - Fine guarnizione torta 4
Pasticcere 3 - Fine decorazione torta 2
Pasticcere 3 - Inizio decorazione torta 3
Pasticcere 2 - Inizio guarnizione torta 5
Torta 2 finita
Pasticcere 2 - Fine guarnizione torta 5
Pasticcere 3 - Fine decorazione torta 3
Pasticcere 3 - Inizio decorazione torta 4
Torta 3 finita
Pasticcere 3 - Fine decorazione torta 4
Pasticcere 3 - Inizio decorazione torta 5
Torta 4 finita
Pasticcere 3 - Fine decorazione torta 5
Torta 5 finita
Produzione di tutte le torte finita!

*/
