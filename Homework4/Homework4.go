/*
    Scrivere un programma in Go che simuli un'attività di trading di valute in un mercato fittizio.

    Il programma deve simulare usando la concorrenza tre coppie di valute: EUR/USD, GBP/USD e JPY/USD, e simulare le operazioni di acquisto e vendita in parallelo.

    Creare una funzione "simulateMarketData" che simuli il prezzo delle coppie di valute e invii i dati simulati su un canale. I prezzi vengono generati e inviati sul canale corrispondente ogni secondo. In particolare:
    • Il prezzo della coppia EUR/USD varia casualmente tra 1.0 e 1.5.
    • Il prezzo della coppia GBP/USD varia casualmente tra 1.0 e 1.5.
    • Il prezzo della coppia JPY/USD varia casualmente tra 0.006 e 0.009.

    Creare una funzione "selectPair" che usa una "select" per gestire le operazioni di vendita e acquisto in base alle condizioni specificate. In particolare:
    • Se il prezzo di EUR/USD supera 1.20, deve vendere EUR/USD. Simulare la vendita con un tempo di 4 secondi, cioè inserire un delay di 4 secondi prima di confermare la vendita.
    • Se il prezzo di GBP/USD scende sotto 1.35, deve acquistare GBP/USD. Simulare l'acquisto con un tempo di 3 secondi, cioè inserire un delay di 3 secondi prima di confermare l'acquisto.
    • Se il prezzo di JPY/USD scende sotto 0.0085, deve acquistare JPY/USD. Simulare l'acquisto con un tempo di 3 secondi, cioè inserire un delay di 3 secondi prima di confermare l'acquisto.

    Il programma deve eseguire il ciclo di trading per un minuto e alla fine del ciclo deve terminare

*/

package main

import (
    "fmt"   // Standard package per la stampa su console 	
    "math/rand" // Package rand per la generazione di numeri casuali
    "time"  // Package time per la gestione degli intervalli di tempo
    "sync"  // Package sync per la sincronizzazione con i mutex
)

// Definisco delle stringhe per identificare le varie coppie di valute
type CurrencyPair string
const (
    EURUSD CurrencyPair = "EUR/USD"
    GBPUSD CurrencyPair = "GBP/USD"
    JPYUSD CurrencyPair = "JPY/USD"
)

/* **** FUNZIONE PER GENERARE I DATI DI MERCATO ****
 * Ogni secondo viene generato un nuovo valore per ogni valuta
 * I valori sono compresi tra un minimo e un massimo specificati per ogni valuta
 * I valori vengono aggiornati sul canale market
*/
func simulateMarketData(market chan map[CurrencyPair]float64) {
    for {
        data := <- market
        data[EURUSD] = rand.Float64()*0.5 + 1.0
        data[GBPUSD] = rand.Float64()*0.5 + 1.0
        data[JPYUSD] = rand.Float64()*0.003 + 0.006
        market <- data
        fmt.Printf("Aggiorna borsa --> EUR/USD: %.3f GBP/USD: %.3f JPY/USD: %.3f\n", data[EURUSD], data[GBPUSD], data[JPYUSD])
        time.Sleep(time.Second)
    }
}

/* **** FUNZIONE PER SELEZIONARE LE COPPIE DI VALUTE ****


*/
func selectPair(marketData chan map[CurrencyPair]float64, wg *sync.Mutex){
    select{
    case data := <-marketData:
        marketData <- data
        wg.Lock()
        if data[EURUSD] > 1.20 {
            fmt.Printf("Sto vendendo %s a %.3f ...\n", EURUSD, data[EURUSD])
            time.Sleep(4 * time.Second)
            fmt.Printf("Transazione %s confermata\n", EURUSD)
        }
        wg.Unlock()
    case data := <-marketData:
        marketData <- data
        wg.Lock()
        if data[GBPUSD] < 1.35 {
            fmt.Printf("Sto comprando %s a %.3f ...\n", GBPUSD, data[GBPUSD])
            time.Sleep(3 * time.Second)
            fmt.Printf("Transazione %s confermata\n", GBPUSD)
        }
        wg.Unlock()
    case data := <-marketData:
        marketData <- data
        wg.Lock()
        if data[JPYUSD] < 0.0085 {
            fmt.Printf("Sto comprando %s a %.3f ...\n", JPYUSD, data[JPYUSD])
            time.Sleep(3 * time.Second)
            fmt.Printf("Transazione %s confermata\n", JPYUSD)
        }
        wg.Unlock()
    }
}

func main(){
    rand.Seed(time.Now().UnixNano()) // Inizializzo il generatore di numeri casuali

    /* **** STRUMENTI PER LA SINCRONIZZAZIONE ****
     * Viene utilizzato un unico canale contenente le varie valute di scambio
     * Viene usato un semplice mutex per la sincronizzazione delle transazioni
    */
    var wg sync.Mutex

    marketData := make(chan map[CurrencyPair]float64, 1)
    marketData <- map[CurrencyPair] float64{} 
    
    // **** INIZIO DEL CICLO DI TRADING ****
    fmt.Println("Starting trading cycle")
    go simulateMarketData(marketData) // Avvio la simulazione del mercato ogni secondo
    
    time.Sleep(time.Second) // Attendo un secondo per far partire la simulazione in modo da avere i primi dati disponibili del mercato
    start := time.Now()
    for time.Since(start) < time.Minute{
        go selectPair(marketData, &wg)
    }

    fmt.Println("Trading cycle ended")
}
