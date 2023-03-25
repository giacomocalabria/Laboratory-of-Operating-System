package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	Nome string
}

type Veicolo struct {
	Tipo string
}

func noleggia(c Cliente, wg *sync.WaitGroup, veicoli []Veicolo, mutex *sync.Mutex, noleggiati map[string]int) {
	defer wg.Done()

	// scegliere un veicolo a caso tra quelli disponibili
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(veicoli))
	v := veicoli[idx]

	// aggiornare il contatore dei veicoli noleggiati
	mutex.Lock()
	noleggiati[v.Tipo]++
	mutex.Unlock()

	// stampare il messaggio di noleggio
	fmt.Printf("%s ha noleggiato un veicolo %s\n", c.Nome, v.Tipo)
}

func stampa(noleggiati map[string]int) {
	fmt.Printf("Berline noleggiate: %d\n", noleggiati["Berlina"])
	fmt.Printf("SUV noleggiati: %d\n", noleggiati["SUV"])
	fmt.Printf("Station Wagon noleggiate: %d\n", noleggiati["Station Wagon"])
}

func main() {
	clienti := []Cliente{
		{"Mario"},
		{"Luigi"},
		{"Carla"},
		{"Giovanni"},
		{"Maria"},
		{"Antonio"},
		{"Francesca"},
		{"Giuseppe"},
		{"Paolo"},
		{"Anna"},
	}

	veicoli := []Veicolo{
		{"Berlina"},
		{"SUV"},
		{"Station Wagon"},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	noleggiati := make(map[string]int)

	for _, c := range clienti {
		wg.Add(1)
		go noleggia(c, &wg, veicoli, &mutex, noleggiati)
	}

	wg.Wait()
	stampa(noleggiati)
}
