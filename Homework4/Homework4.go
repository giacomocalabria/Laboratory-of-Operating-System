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
