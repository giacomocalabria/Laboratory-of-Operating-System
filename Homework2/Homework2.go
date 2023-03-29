/* Questo programma simula un'agenzia di noleggi d'auto, gestisce le prenotazioni di 10 clienti.

Un cliente noleggia un veicolo. 
I veicoli disponibili sono: Berlina, SUV o Station Wagon.

Ogni cliente può noleggiare un veicolo contemporaneamente ad altri.

*/

package main

import (
	"fmt"
)

// Creo la struttura Cliente con il campo "nome"
type Cliente struct{
	nome string
}

// Creo la struttura Veicolo con il campo "tipo"
type Veicolo struct{
	tipo string
}

//Creare  la  function  "noleggia"  che  prende  come  input  un cliente  e  che  prenota uno a caso tra i veicoli. Questa function deve anche stampare che il cliente x ha noleggiato il veicolo y.
func noleggia(c Cliente){
	fmt.Printf("Ciao")
}

//Creare  una  function  "stampa"  che,  alla  fine  del  processo,  stampa  il  numero  di Berline, SUV e Station Wagon noleggiati.
func stampa(){
	fmt.Printf("Veicoli tipo Berlina noleggiati: %d\n", )
	fmt.Printf("Veicoli tipo SUV noleggiati: %d\n", )
	fmt.Printf("Veicoli tipo Station Wagon noleggiati: %d\n", )
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
	
	stampa()
}