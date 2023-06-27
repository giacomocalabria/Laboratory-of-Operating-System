/* Questo programma simula un'agenzia di noleggi d'auto, gestisce le prenotazioni di 10 clienti.

Un cliente noleggia un veicolo. 
I veicoli disponibili sono: Berlina, SUV o Station Wagon.

Ogni cliente può noleggiare un veicolo contemporaneamente ad altri.


• Creare la struttura Cliente con il campo "nome"
• Creare la struttura Veicolo con il campo "tipo"

• Creare la function "noleggia" che prende come input un cliente e che prenota uno a caso tra i veicoli. Questa function deve anche stampare che il cliente x ha noleggiato il veicolo y.

• Creare una function "stampa" che, alla fine del processo, stampa il numero di
Berline, SUV e Station Wagon noleggiati.

• Ogni cliente può noleggiare un veicolo contemporaneamente ad altri.

Si noti che si possono creare ulteriori funzioni per risolvere il problema, oltre alle due obbligatorie, descritte sopra.

*/

package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

// Creo la struttura Cliente con il campo "nome"
type Cliente struct{
	nome string
}

// Creo la struttura Veicolo con il campo "tipo"
type Veicolo struct{
	tipo string
}

func increaseNumber(conta chan int) {
	num := <-conta
	num++
	conta <- num
}

//Creare  la  function  "noleggia"  che  prende  come  input  un cliente  e  che  prenota uno a caso tra i veicoli. Questa function deve anche stampare che il cliente x ha noleggiato il veicolo y.
func noleggia(c Cliente, veicoli []Veicolo, conta1 chan int, conta2 chan int, conta3 chan int, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(veicoli))

	switch i {
	case 0:
		increaseNumber(conta1)
	case 1:
		increaseNumber(conta2)
	case 2:
		increaseNumber(conta3)
	}
	veicolo := veicoli[i]
	fmt.Printf("Il Cliente %s ha noleggiato il veicolo %s.\n", c.nome, veicolo.tipo)

	defer wg.Done()
}

//Creare  una  function  "stampa"  che,  alla  fine  del  processo,  stampa  il  numero  di Berline, SUV e Station Wagon noleggiati.
func stampa(conta1 chan int, conta2 chan int, conta3 chan int) {
	num1 := <- conta1
	num2 := <- conta2
	num3 := <- conta3

	fmt.Println("Numero di veicoli noleggiati per tipo:")
	fmt.Printf("Berline-> %d\n", num1)
	fmt.Printf("SUV-> %d\n", num2)
	fmt.Printf("Station Wagon-> %d\n", num3)
}

func main(){
	clienti := []Cliente{
		{"A"},
		{"B"},
		{"C"},
		{"D"},
		{"E"},
		{"F"},
		{"G"},
		{"H"},
		{"I"},
		{"L"},
	}

	veicoli := []Veicolo{
		{"Berlina"},
		{"SUV"},
		{"Station Wagon"},
	}

	conta1 := make(chan int, 1)
	conta2 := make(chan int, 1)
	conta3 := make(chan int, 1)
	conta1 <- 0
	conta2 <- 0
	conta3 <- 0
	
	var wg sync.WaitGroup // WaitGroup per la sincronizzazione

	for _, cliente := range clienti{
		wg.Add(1)
		go noleggia(cliente, veicoli, conta1, conta2, conta3, &wg)
	}

	wg.Wait()

	stampa(conta1, conta2, conta3)
}