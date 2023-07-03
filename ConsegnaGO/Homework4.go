/*
    Author: Giacomo Calabria
	Homework 4
    
    Questo programma simula un'attività di trading di valute in un mercato fittizio, esegue un ciclo di trading per un minuto e alla fine del ciclo termina.

    Il programma deve simulare usando la concorrenza tre coppie di valute: EUR/USD, GBP/USD e JPY/USD, e simulare le operazioni di acquisto e vendita in parallelo.

    La funzione "simulateMarketData" simula il prezzo delle coppie di valute e invia i dati simulati su un canale. I prezzi vengono generati e inviati sul canale corrispondente ogni secondo. In particolare:
    • Il prezzo della coppia EUR/USD varia casualmente tra 1.0 e 1.5.
    • Il prezzo della coppia GBP/USD varia casualmente tra 1.0 e 1.5.
    • Il prezzo della coppia JPY/USD varia casualmente tra 0.006 e 0.009.

    La funzione "selectPair" che usa una "select" per gestire le operazioni di vendita e acquisto in base alle condizioni specificate. In particolare:
    • Se il prezzo di EUR/USD supera 1.20, deve vendere EUR/USD dopo un delay di 4s per la conferma.
    • Se il prezzo di GBP/USD scende sotto 1.35, deve acquistare GBP/USD dopo un delay di 3s per la conferma.
    • Se il prezzo di JPY/USD scende sotto 0.0085, deve acquistare JPY/USD dopo un delay di 3s per la conferma.

    Utilizza un Channel per i valori del mercato e un mutex per la sincronizzazione delle transazioni.
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
func simulateMarketData(market chan map[CurrencyPair]float64, tradingInCorso *bool) {
    for *tradingInCorso{
        data := <- market
        data[EURUSD] = rand.Float64()*0.5 + 1.0
        data[GBPUSD] = rand.Float64()*0.5 + 1.0
        data[JPYUSD] = rand.Float64()*0.003 + 0.006
        market <- data
        //fmt.Printf("Aggiorna borsa --> EUR/USD: %.3f GBP/USD: %.3f JPY/USD: %.3f\n", data[EURUSD], data[GBPUSD], data[JPYUSD])
        time.Sleep(time.Second)
    }
}

// **** FUNZIONE PER SELEZIONARE LE COPPIE DI VALUTE ****
func selectPair(marketData chan map[CurrencyPair]float64,tradingInCorso *bool, wg *sync.Mutex){
    for *tradingInCorso{
        wg.Lock() // Aspetto che il mercato sia libero per poter effettuare la transazione
        
        // Seleziono la coppia di valute in base al prezzo
        select{
        case data := <-marketData:
            marketData <- data
            if data[EURUSD] > 1.20 {
                fmt.Printf("Sto vendendo %s a %.3f ...\n", EURUSD, data[EURUSD])
                time.Sleep(4 * time.Second)
                fmt.Printf("Transazione %s confermata\n", EURUSD)
            }
        case data := <-marketData:
            marketData <- data
            if data[GBPUSD] < 1.35 {
                fmt.Printf("Sto comprando %s a %.3f ...\n", GBPUSD, data[GBPUSD])
                time.Sleep(3 * time.Second)
                fmt.Printf("Transazione %s confermata\n", GBPUSD)
            }
        case data := <-marketData:
            marketData <- data
            if data[JPYUSD] < 0.0085 {
                fmt.Printf("Sto comprando %s a %.3f ...\n", JPYUSD, data[JPYUSD])
                time.Sleep(3 * time.Second)
                fmt.Printf("Transazione %s confermata\n", JPYUSD)
            }
            
        }
        wg.Unlock() // Rilascio il mercato
    }
}

func main(){
    rand.Seed(time.Now().UnixNano()) // Inizializzo il generatore di numeri casuali

    /* **** STRUMENTI PER LA SINCRONIZZAZIONE ****
     * Viene utilizzato un unico canale contenente le varie valute di scambio
     * Viene usato un semplice mutex per la sincronizzazione delle transazioni
     * Viene usata una variabile booleana per la temporiazzaione del ciclo di trading
    */
    var wg sync.Mutex
    var tradingInCorso bool = true 

    marketData := make(chan map[CurrencyPair]float64, 1)
    marketData <- map[CurrencyPair] float64{}
    
    // **** INIZIO DEL CICLO DI TRADING ****
    fmt.Println("Starting trading cycle")

    go simulateMarketData(marketData, &tradingInCorso) // Avvio la simulazione del mercato ogni secondo
    time.Sleep(time.Millisecond) // Attendo che il canale sia popolato dai dati di mercato

    go selectPair(marketData, &tradingInCorso, &wg) // Avvio la selezione delle coppie di valute

    // **** GESTIONE FINE CICLO ****
    time.Sleep(time.Minute) // Attendo un minuto prima di terminare il ciclo di trading
    for !wg.TryLock() {} // Attendo che tutte le transazioni siano concluse
    tradingInCorso = false // Fermo la selezione delle coppie di valute

    fmt.Println("Trading cycle ended")
}

/* OUTPUT ESEMPIO:

Starting trading cycle
Sto vendendo EUR/USD a 1.326 ...
Transazione EUR/USD confermata
Sto comprando JPY/USD a 0.007 ...
Transazione JPY/USD confermata
Sto comprando JPY/USD a 0.008 ...
Transazione JPY/USD confermata
Sto comprando JPY/USD a 0.007 ...
Transazione JPY/USD confermata
Sto comprando JPY/USD a 0.007 ...
Transazione JPY/USD confermata
Sto comprando GBP/USD a 1.133 ...
Transazione GBP/USD confermata
Sto comprando JPY/USD a 0.007 ...
Transazione JPY/USD confermata
Sto vendendo EUR/USD a 1.304 ...
Transazione EUR/USD confermata
Sto vendendo EUR/USD a 1.306 ...
Transazione EUR/USD confermata
Sto comprando GBP/USD a 1.045 ...
Transazione GBP/USD confermata
Sto comprando JPY/USD a 0.006 ...
Transazione JPY/USD confermata
Sto vendendo EUR/USD a 1.416 ...
Transazione EUR/USD confermata
Sto comprando GBP/USD a 1.310 ...
Transazione GBP/USD confermata
Sto comprando JPY/USD a 0.006 ...
Transazione JPY/USD confermata
Sto vendendo EUR/USD a 1.337 ...
Transazione EUR/USD confermata
Sto comprando GBP/USD a 1.020 ...
Transazione GBP/USD confermata
Sto comprando GBP/USD a 1.206 ...
Transazione GBP/USD confermata
Sto comprando JPY/USD a 0.007 ...
Transazione JPY/USD confermata
Sto comprando JPY/USD a 0.006 ...
Transazione JPY/USD confermata
Trading cycle ended

*/
