/*

Scrivere un programma in Go che simuli la produzione di 5 torte da parte di 3
pasticceri. La produzione di ogni torta richiede 3 fasi che devono avvenire in ordine:
prima la torta viene cucinata, poi guarnita e infine decorata.

Il primo pasticcere si occupa solo di cucinare le torte e ci mette 1 secondo per ogni
torta. Questo pasticcere ha a disposizione 2 spazi per appoggiare le torte una volta
che ha finito di cucinarle. Se ci sono spazi liberi, può iniziare a cucinare la torta
successiva senza aspettare che il secondo pasticcere si liberi per guarnire quella
appena cucinata. 

Il secondo pasticcere si occupa solo di guarnire le torte e ci mette 4
secondi per ogni torta. Anche questo pasticcere ha a disposizione 2 spazi per
appoggiare le torte una volta che ha finito di guarnirle. 

Il terzo pasticcere si occupa solo di decorare le torte e ci mette 8 secondi per ogni torta.


I tre pasticceri lavorano contemporaneamente.

*/

package main

import (
	"fmt"
	"time"
)

func pasticcere1(spazi1 chan int, TorteCucinate chan int){
	for i := 1; i <= 5; i++{
		spazi1 <- i
		fmt.Printf("Pasticcere 1 - Inizio cottura torta %d\n", i)
		time.Sleep(1 * time.Second)
		fmt.Printf("Pasticcere 1 - Fine cottura torta %d\n", i)
		TorteCucinate <- i
	}
	fmt.Println("Cottura torte completata!")
}

func pasticcere2(spazi1 chan int, spazi2 chan bool, TorteCucinate chan int, TorteFarcite chan int){
	for i := 1; i <= 5; i++{
		torta := <- TorteCucinate
		spazi2 <- true
		<- spazi1
		fmt.Printf("Pasticcere 2 - Inizio guarnizione torta %d\n", torta)
		time.Sleep(4 * time.Second)
		fmt.Printf("Pasticcere 2 - Fine guarnizione torta %d\n", torta)
		TorteFarcite <- torta
	}
}

func pasticcere3(spazi2 chan bool, TorteFarcite chan int, TorteFinite chan int){
	for i := 1; i <= 5; i++{
		torta := <-TorteFarcite
		<- spazi2
		fmt.Printf("Pasticcere 3 - Inizio decorazione torta %d\n", torta)
		time.Sleep(8 * time.Second)
		fmt.Printf("Pasticcere 3 - Fine decorazione torta %d\n", torta)
		TorteFinite <- torta
	}
}

func main() {
	// Canale per gli spazi liberi per appoggiare le torne una volta che hanno finito di cucinarle. Se ci sono spazi liberi, può iniziare a cucinare la torta successiva senza aspettare che il secondo pasticcere si liberi per guarnire quella appena cucinata.
	spazi1 := make(chan int, 2)
	// Canale per sapere se c'è una torna da cucinare
	TorteCucinate := make(chan int, 5)

	// Canale per gli spazi libero per appoggiare le torte una volta che ha finito di guarnirle. 
	spazi2 := make(chan bool, 2)

	// Canale per sapere se c'è una torta da farcire
	TorteFarcite := make(chan int, 5)

	// Canale per sapere se una torta è finita
	TorteFinite := make(chan int, 5)

	fmt.Println("Produzione torte iniziata!")

	go pasticcere1(spazi1, TorteCucinate)

	go pasticcere2(spazi1, spazi2, TorteCucinate, TorteFarcite)

	go pasticcere3(spazi2, TorteFarcite, TorteFinite)

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
