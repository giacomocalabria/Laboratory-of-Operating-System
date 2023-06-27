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
    "fmt"
    "math/rand"
    "time"
)

type CurrencyPair string

const (
    EURUSD CurrencyPair = "EUR/USD"
    GBPUSD CurrencyPair = "GBP/USD"
    JPYUSD CurrencyPair = "JPY/USD"
)

type TradeAction string

const (
    Buy  TradeAction = "buy"
    Sell TradeAction = "sell"
)

type Trade struct {
    Pair   CurrencyPair
    Action TradeAction
}

func simulateMarketData(ch chan<- map[CurrencyPair]float64) {
    for {
        data := make(map[CurrencyPair]float64)
        data[EURUSD] = rand.Float64()*0.5 + 1.0
        data[GBPUSD] = rand.Float64()*0.5 + 1.0
        data[JPYUSD] = rand.Float64()*0.003 + 0.006
        ch <- data
        time.Sleep(time.Second)
    }
}

func selectPair(ch <-chan map[CurrencyPair]float64, trades chan<- Trade) {
    for {
        select {
        case data := <-ch:
            if data[EURUSD] > 1.20 {
                trades <- Trade{Pair: EURUSD, Action: Sell}
                time.Sleep(4 * time.Second)
                fmt.Printf("Sold %s at %.4f\n", EURUSD, data[EURUSD])
            } else if data[GBPUSD] < 1.35 {
                trades <- Trade{Pair: GBPUSD, Action: Buy}
                time.Sleep(3 * time.Second)
                fmt.Printf("Bought %s at %.4f\n", GBPUSD, data[GBPUSD])
            } else if data[JPYUSD] < 0.0085 {
                trades <- Trade{Pair: JPYUSD, Action: Buy}
                time.Sleep(3 * time.Second)
                fmt.Printf("Bought %s at %.4f\n", JPYUSD, data[JPYUSD])
            }
        case <-time.After(time.Minute):
            fmt.Println("Trading cycle completed")
            return
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    marketData := make(chan map[CurrencyPair]float64)
    trades := make(chan Trade)

    go simulateMarketData(marketData)
    go selectPair(marketData, trades)

    <-time.After(time.Minute)
}
